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

type hasher interface {
	Hash(src []byte) (res, salt []byte, err error)
}

type keepService struct {
	card    cardStorage
	creds   credentialsStorage
	attachs attachmentsStorage
	hasher  hasher
}

func NewService(cardStorage cardStorage, credsStorage credentialsStorage, attachsStorage attachmentsStorage, hasher hasher) *keepService {
	return &keepService{
		card:    cardStorage,
		creds:   credsStorage,
		attachs: attachsStorage,
	}
}
