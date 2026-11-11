package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"

	"github.com/dDan4a/go-template/internal/config"
	httpCont "github.com/dDan4a/go-template/internal/controller/http"
	notePg "github.com/dDan4a/go-template/internal/repository/note/pg"
	noteSvc "github.com/dDan4a/go-template/internal/service/note"
	"github.com/dDan4a/go-template/internal/version"
	"github.com/dDan4a/go-template/migrations"
)

const shutdownTimeout = 10 * time.Second

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return run(cmd.Context())
	},
}

func NewHandler(db *pgxpool.Pool) http.Handler {
	noteRepository := notePg.New(db)
	noteService := noteSvc.New(noteRepository)

	return httpCont.New(noteService).Routes()
}

func run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("connect postgres: %w", err)
	}
	defer db.Close()

	if _, err := db.Exec(ctx, migrations.Schema); err != nil {
		return fmt.Errorf("apply schema: %w", err)
	}

	srv := &http.Server{Addr: cfg.HTTPAddr, Handler: NewHandler(db), ReadHeaderTimeout: 5 * time.Second}

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), shutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			slog.Error("shutdown http server", slog.Any("error", err))
		}
	}()

	slog.Info("serving http", slog.String("addr", cfg.HTTPAddr), slog.String("commit", version.GitCommit))

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
