package keep

import (
	"context"
	"errors"

	"github.com/arxon31/gophkeep/internal/models"
	"github.com/arxon31/gophkeep/internal/models/requests"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"

	"github.com/arxon31/yapr-proto/pkg/gophkeep"
)

var ErrUnknownRequest = errors.New("unknown request to keep")

type textStorage interface {
	SaveCredentials(ctx context.Context, dto requests.SaveCredentialsDTO) error
	SaveBankCredentials(ctx context.Context, dto requests.SaveBankCredentialsDTO) error
	SaveFileS3URL(ctx context.Context, dto requests.SaveFileS3URLDTO) error
}

type fileStorage interface {
	SaveFile(ctx context.Context, file *models.FileDTO) (url string, err error)
}

type keepService struct {
	text textStorage
	file fileStorage
}

func NewService(textStorage textStorage, fileStorage fileStorage) *keepService {
	return &keepService{text: textStorage, file: fileStorage}
}

func (s *keepService) Keep(ctx context.Context, record interface{}) error {
	switch rec := record.(type) {
	case *gophkeep.SaveCredentialsRequest:
		return s.keepCredentials(ctx, rec)
	case *gophkeep.SaveBankCredentialsRequest:
		return s.keepBankCredentials(ctx, rec)
	case *models.FileDTO:
		return s.keepFile(ctx, rec)
	default:
		return ErrUnknownRequest
	}

}

func (s *keepService) keepCredentials(ctx context.Context, creds *gophkeep.SaveCredentialsRequest) error {
	user, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	credsDTO := requests.SaveCredentialsDTO{
		User:     user,
		Meta:     creds.Meta.GetMeta(),
		UserName: creds.GetUsername(),
		Password: creds.GetPassword(),
	}

	err = credsDTO.Validate()
	if err != nil {
		return err
	}

	return s.text.SaveCredentials(ctx, credsDTO)
}

func (s *keepService) keepBankCredentials(ctx context.Context, bankCreds *gophkeep.SaveBankCredentialsRequest) error {
	user, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	bankCredsDTO := requests.SaveBankCredentialsDTO{
		User:       user,
		Meta:       bankCreds.Meta.GetMeta(),
		CardNumber: bankCreds.GetCardNumber(),
		Owner:      bankCreds.GetOwner(),
		CVV:        bankCreds.GetCvv(),
	}

	err = bankCredsDTO.Validate()
	if err != nil {
		return err
	}

	return s.text.SaveBankCredentials(ctx, bankCredsDTO)
}

func (s *keepService) keepFile(ctx context.Context, file *models.FileDTO) error {
	user, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	file.User = user

	err = file.Validate()
	if err != nil {
		return err
	}

	url, err := s.file.SaveFile(ctx, file)
	if err != nil {
		return err
	}

	s3URLDTO := requests.SaveFileS3URLDTO{
		User: user,
		Meta: file.Meta,
		URL:  url,
	}

	err = s3URLDTO.Validate()
	if err != nil {
		return err
	}

	return s.text.SaveFileS3URL(ctx, s3URLDTO)

}
