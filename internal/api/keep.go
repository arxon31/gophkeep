package api

import (
	"context"
	"errors"
	"github.com/arxon31/gophkeep/internal/converter"
	"github.com/arxon31/gophkeep/internal/model/attachment"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"
	"google.golang.org/grpc"
	"io"
)

const (
	success = true
	failure = false
)

type fileInfo struct {
	name string
	meta string
}

func (s *server) SaveCard(ctx context.Context, card *gophkeep.SaveCardRequest) (*gophkeep.SaveStatus, error) {
	serviceCard, meta := converter.CardToService(card)

	err := s.keep.KeepCard(ctx, serviceCard, meta)
	if err != nil {
		keepErrorMessage := err.Error()
		return &gophkeep.SaveStatus{Success: failure, Error: &keepErrorMessage}, nil
	}

	return &gophkeep.SaveStatus{Success: success}, nil

}

func (s *server) SaveCredentials(ctx context.Context, credentials *gophkeep.SaveCredentialsRequest) (*gophkeep.SaveStatus, error) {
	serviceCreds, meta := converter.CredentialsToService(credentials)

	err := s.keep.KeepCredentials(ctx, serviceCreds, meta)
	if err != nil {
		keepErrorMessage := err.Error()
		return &gophkeep.SaveStatus{Success: failure, Error: &keepErrorMessage}, nil
	}

	return &gophkeep.SaveStatus{Success: success}, nil
}

func (s *server) SaveAttachment(ctx context.Context, attach *gophkeep.SaveAttachmentRequest) (*gophkeep.SaveAttachmentResponse, error) {
	sessionUUID := s.session.Create(fileInfo{name: attach.GetName(), meta: attach.Meta.GetMeta()})

	return &gophkeep.SaveAttachmentResponse{SessionHash: []byte(sessionUUID)}, nil

}

func (s *server) StartSaveFileStream(srv grpc.ClientStreamingServer[gophkeep.Chunk, gophkeep.SaveStatus]) error {
	sessionUUID, err := ctxfuncs.GetSessionHashFromContext(srv.Context())
	if err != nil {
		errMsg := ErrSessionNotFound.Error()
		_ = srv.SendAndClose(&gophkeep.SaveStatus{Success: failure, Error: &errMsg})
	}

	iInfo, ok := s.session.Info(sessionUUID)
	if !ok {
		errMsg := ErrSessionNotFound.Error()
		_ = srv.SendAndClose(&gophkeep.SaveStatus{Success: failure, Error: &errMsg})
	}

	info, ok := iInfo.(*fileInfo)
	if !ok {
		errMsg := ErrSessionNotFound.Error()
		_ = srv.SendAndClose(&gophkeep.SaveStatus{Success: failure, Error: &errMsg})
	}

	file := &attachment.Attachment{
		Name:    info.name,
		Content: make([]byte, 0, chunkSize),
	}

	attachMeta := meta.Meta(info.meta)

RECEIVE:
	for {
		chunk, err := srv.Recv()

		switch {
		case errors.Is(err, io.EOF):
			break RECEIVE
		default:
			errMsg := ErrSomethingWentWrong.Error()
			_ = srv.SendAndClose(&gophkeep.SaveStatus{Success: failure, Error: &errMsg})

		}

		file.Content = append(file.Content, chunk.GetContent()...)
	}

	err = s.keep.KeepAttachment(srv.Context(), file, attachMeta)
	if err != nil {
		keepErrorMessage := err.Error()
		_ = srv.SendAndClose(&gophkeep.SaveStatus{Success: failure, Error: &keepErrorMessage})
	}

	s.session.Delete(sessionUUID)

	return nil
}
