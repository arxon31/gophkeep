package main

import (
	"context"
	"github.com/arxon31/gophkeep/internal/api"
	"github.com/arxon31/gophkeep/internal/config"
	"github.com/arxon31/gophkeep/internal/repository/attachment"
	"github.com/arxon31/gophkeep/internal/repository/card"
	"github.com/arxon31/gophkeep/internal/repository/credentials"
	"github.com/arxon31/gophkeep/internal/service/keep"
	"github.com/arxon31/gophkeep/internal/service/session"
	"github.com/arxon31/gophkeep/internal/service/sync"
	"github.com/arxon31/gophkeep/pkg/encrypt"
	. "github.com/arxon31/gophkeep/pkg/logger"
	"github.com/arxon31/gophkeep/pkg/minioconn"
	"github.com/arxon31/gophkeep/pkg/mongoconn"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
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

	srv := errgroup.Group{}

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	mongoClient, err := mongoconn.New(ctx, cfg.Mongo.URI)
	if err != nil {
		return err
	}
	defer mongoClient.Disconnect(ctx)

	mongoDB := mongoClient.Database(cfg.Mongo.DBName)

	minioClient, err := minioconn.New(cfg.S3.URI, cfg.S3.User, cfg.S3.Password)
	if err != nil {
		return err
	}

	encryptor, err := encrypt.New([]byte(cfg.Secrets.CryptoKey))
	if err != nil {
		return err
	}

	cardRepo := card.New(mongoDB)
	credentialsRepo := credentials.New(mongoDB)
	attachmentsRepo := attachment.New(minioClient, mongoDB)

	keepService := keep.NewService(cardRepo, credentialsRepo, attachmentsRepo, encryptor)
	syncService := sync.NewService(cardRepo, credentialsRepo, attachmentsRepo, encryptor)
	sessionService := session.NewService()

	server := api.NewServer(keepService, syncService, sessionService)

	srv.Go(func() error {
		return server.Run(ctx)
	})

	select {
	case <-ctx.Done():

	}

	return srv.Wait()
}
