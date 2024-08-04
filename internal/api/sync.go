package api

import (
	"context"
	"github.com/arxon31/gophkeep/internal/converter"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"
	"google.golang.org/grpc"
)

func (s *server) GetCredentials(ctx context.Context, req *gophkeep.GetByMetaRequest) (*gophkeep.GetCredentialsResponse, error) {
	metaFromReq := req.Meta.GetMeta()

	serviceMeta := meta.Meta(metaFromReq)

	creds, err := s.sync.SyncCredentials(ctx, &serviceMeta)
	if err != nil {
		return nil, err
	}

	return converter.CredentialsToProto(creds), nil
}
func (s *server) GetCard(ctx context.Context, req *gophkeep.GetByMetaRequest) (*gophkeep.GetCardResponse, error) {
	metaFromReq := req.Meta.GetMeta()

	serviceMeta := meta.Meta(metaFromReq)

	card, err := s.sync.SyncCard(ctx, &serviceMeta)
	if err != nil {
		return nil, err
	}

	return converter.CardToProto(card), nil

}
func (s *server) GetAttachment(ctx context.Context, req *gophkeep.GetByMetaRequest) (*gophkeep.GetAttachmentResponse, error) {
	metaFromReq := req.Meta.GetMeta()

	serviceMeta := meta.Meta(metaFromReq)

	sessionUUID := s.session.Create(serviceMeta)

	return &gophkeep.GetAttachmentResponse{SessionHash: []byte(sessionUUID)}, nil
}

// TODO: IMPLEMENT ME
func (s *server) StartGetFileStream(*gophkeep.StartGetFileStreamRequest, grpc.ServerStreamingServer[gophkeep.Chunk]) error {
	return nil
}
