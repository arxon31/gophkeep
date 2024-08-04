package keep

import (
	"context"

	attachsmodel "github.com/arxon31/gophkeep/internal/repository/attachment/dto"
	cardmodel "github.com/arxon31/gophkeep/internal/repository/card/dto"
	credsmodel "github.com/arxon31/gophkeep/internal/repository/credentials/dto"
)

type cardStorage interface {
	SaveCard(ctx context.Context, card *cardmodel.Card) error
}

type credentialsStorage interface {
	SaveCredentials(ctx context.Context, creds *credsmodel.Credentials) error
}

type attachmentsStorage interface {
	SaveAttachment(ctx context.Context, attachmentInfo *attachsmodel.Attachment) error
}

type encryptor interface {
	Encrypt(src []byte) ([]byte, error)
}

type keepService struct {
	card      cardStorage
	creds     credentialsStorage
	attachs   attachmentsStorage
	encryptor encryptor
}

func NewService(cardStorage cardStorage, credsStorage credentialsStorage, attachsStorage attachmentsStorage, encryptor encryptor) *keepService {
	return &keepService{
		card:      cardStorage,
		creds:     credsStorage,
		attachs:   attachsStorage,
		encryptor: encryptor,
	}
}
