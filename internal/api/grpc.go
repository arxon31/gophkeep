package api

import (
	"context"
	"github.com/arxon31/gophkeep/internal/model/attachment"
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"
	"google.golang.org/grpc"
	"net"
)

//go:generate moq -out keep_moq_test.go . keepService
type keepService interface {
	KeepAttachment(ctx context.Context, attach *attachment.Attachment, attachMeta meta.Meta) error
	KeepCard(ctx context.Context, card *card.Card, cardMeta meta.Meta) error
	KeepCredentials(ctx context.Context, creds *credentials.Credentials, credsMeta meta.Meta) error
}

//go:generate moq -out sync_moq_test.go . syncService
type syncService interface {
	SyncAttachment(ctx context.Context, req *meta.Meta) (*attachment.Attachment, error)
	SyncCard(ctx context.Context, req *meta.Meta) (resp *card.Card, err error)
	SyncCredentials(ctx context.Context, req *meta.Meta) (resp *credentials.Credentials, err error)
}

//go:generate moq -out session_moq_test.go . sessionService
type sessionService interface {
	Create(info any) (sessionID string)
	Info(sessionID string) (any, bool)
	Delete(sessionID string)
}

const chunkSize = 1024

const serverPort = ":8080"

type server struct {
	gophkeep.UnimplementedGophKeepServer

	keep    keepService
	sync    syncService
	session sessionService
}

func NewServer(keep keepService, sync syncService, session sessionService) *server {
	return &server{
		keep:    keep,
		sync:    sync,
		session: session,
	}
}

func (s *server) Run(ctx context.Context) error {
	srv := grpc.NewServer()

	gophkeep.RegisterGophKeepServer(srv, s)

	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		srv.GracefulStop()
	}()

	return srv.Serve(listener)
}
