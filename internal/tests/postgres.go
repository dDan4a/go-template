package tests

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/dDan4a/go-template/migrations"
)

func Postgres(t *testing.T) *pgxpool.Pool {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Fatal("DATABASE_URL is not set: run docker compose up -d and make e2e")
	}

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatalf("connect postgres: %v", err)
	}

	t.Cleanup(db.Close)

	if _, err := db.Exec(t.Context(), migrations.Schema); err != nil {
		t.Fatalf("apply schema: %v", err)
	}

	return db
}
