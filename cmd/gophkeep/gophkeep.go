package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	. "github.com/arxon31/gophkeep/pkg/logger"
)

var Build string

func main() {
	Logger.Info("starting app", slog.String("build", Build))
	err := run()
	if err != nil {
		Logger.Error("app exited with error", slog.String("error", err.Error()))
		os.Exit(1)
	}
	Logger.Info("app exited without errors")
	os.Exit(0)
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	return nil
}
