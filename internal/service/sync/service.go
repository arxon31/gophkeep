package sync

import (
	"context"
	attachmodel "github.com/arxon31/gophkeep/internal/repository/attachment/dto"
	cardmodel "github.com/arxon31/gophkeep/internal/repository/card/dto"
	credsmodel "github.com/arxon31/gophkeep/internal/repository/credentials/dto"
)

type cardProvider interface {
	GetCard(ctx context.Context, req *cardmodel.GetCard) (*cardmodel.Card, error)
}

type credentialsProvider interface {
	GetCredentials(ctx context.Context, req *credsmodel.GetCredentials) (*credsmodel.Credentials, error)
}

type attachmentsProvider interface {
	GetAttachment(ctx context.Context, req *attachmodel.GetAttachment) (*attachmodel.Attachment, error)
}

type decryptor interface {
	Decrypt([]byte) ([]byte, error)
}

type syncService struct {
	card        cardProvider
	creds       credentialsProvider
	attachments attachmentsProvider
	decryptor   decryptor
}

func NewService(card cardProvider, creds credentialsProvider, attachments attachmentsProvider, decryptor decryptor) *syncService {
	return &syncService{card: card, creds: creds, attachments: attachments, decryptor: decryptor}
}
