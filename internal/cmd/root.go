package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/dDan4a/go-template/internal/cmd/server"
)

var rootCmd = &cobra.Command{SilenceUsage: true, SilenceErrors: true}

func init() {
	rootCmd.AddCommand(server.Cmd)
}

func Execute() {
	if err := execute(); err != nil {
		log.Fatalf("failed to execute root cmd: %s", err)
	}
}

func execute() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	return rootCmd.ExecuteContext(ctx)
}
