package sync

import (
	"context"

	"github.com/arxon31/yapr-proto/pkg/gophkeep"

	"github.com/arxon31/gophkeep/internal/model/requests"
	"github.com/arxon31/gophkeep/internal/model/responses"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
)

type provider interface {
	GetCredentials(ctx context.Context, dto requests.GetByMetaDTO) (resp *responses.GetCredentialsResponseDTO, err error)
	GetBankCredentials(ctx context.Context, dto requests.GetByMetaDTO) (resp *responses.GetBankCredentialsResponseDTO, err error)
	GetFileS3URL(ctx context.Context, dto requests.GetByMetaDTO) (resp *responses.GetS3FileURLDTO, err error)
}

type syncService struct {
	provider provider
}

func NewService(provider provider) *syncService {
	return &syncService{provider: provider}
}

func (s *syncService) SyncCredentials(ctx context.Context, req *gophkeep.GetByMetaRequest) (resp *gophkeep.GetCredentialsResponse, err error) {
	user, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	dto := &requests.GetByMetaDTO{
		User: user,
		Meta: req.Meta.GetMeta(),
	}

	err = dto.Validate()
	if err != nil {
		return nil, err
	}

}

func (s *syncService) SyncBankCredentials(ctx context.Context, req *gophkeep.GetByMetaRequest) (resp *gophkeep.GetBankCredentialsResponse, err error) {

}

func (s *syncService) SyncFile(ctx context.Context, req *gophkeep.GetByMetaRequest) (resp *gophkeep.GetFileResponse, err error) {

}
