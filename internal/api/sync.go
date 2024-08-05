package api

import (
	"context"
	"github.com/arxon31/gophkeep/internal/converter"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
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

	sessionUUID := s.session.Create(&serviceMeta)

	return &gophkeep.GetAttachmentResponse{SessionHash: []byte(sessionUUID)}, nil
}

func (s *server) StartGetFileStream(req *gophkeep.StartGetFileStreamRequest, srv grpc.ServerStreamingServer[gophkeep.Chunk]) error {
	sessionUUID, err := ctxfuncs.GetSessionHashFromContext(srv.Context())
	if err != nil {
		return err
	}

	iMeta, ok := s.session.Info(sessionUUID)
	if !ok {
		return ErrSessionNotFound
	}

	meta, ok := iMeta.(*meta.Meta)
	if !ok {
		return ErrSessionNotFound
	}

	attach, err := s.sync.SyncAttachment(srv.Context(), meta)
	if err != nil {
		return err
	}

	chunks := splitContent(attach.Content, chunkSize)

	for _, chunk := range chunks {
		err = srv.Send(&gophkeep.Chunk{Content: chunk})
		if err != nil {
			return err
		}
	}

	return nil

}

func splitContent(content []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(content); i += chunkSize {
		end := i + chunkSize
		if end > len(content) {
			end = len(content)
		}
		chunks = append(chunks, content[i:end])
	}
	return chunks
}
