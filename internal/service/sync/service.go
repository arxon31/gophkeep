package sync

import (
	"context"
	attachmodel "github.com/arxon31/gophkeep/internal/repository/attachment/dto"
	cardmodel "github.com/arxon31/gophkeep/internal/repository/card/dto"
	credsmodel "github.com/arxon31/gophkeep/internal/repository/credentials/dto"
)

//go:generate moq -out card_moq_test.go . cardProvider
type cardProvider interface {
	GetCard(ctx context.Context, req *cardmodel.GetCard) (*cardmodel.Card, error)
}

//go:generate moq -out creds_moq_test.go . credentialsProvider
type credentialsProvider interface {
	GetCredentials(ctx context.Context, req *credsmodel.GetCredentials) (*credsmodel.Credentials, error)
}

//go:generate moq -out attachs_moq_test.go . attachmentsProvider
type attachmentsProvider interface {
	GetAttachment(ctx context.Context, req *attachmodel.GetAttachment) (*attachmodel.Attachment, error)
}

//go:generate moq -out decryptor_moq_test.go . decryptor
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
