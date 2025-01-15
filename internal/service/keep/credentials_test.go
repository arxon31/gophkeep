package keep

import (
	"context"
	"errors"
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/repository/credentials/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeepService_KeepCredentials(t *testing.T) {
	ctxWithoutUser := context.Background()
	ctxWithUser := ctxfuncs.SetUserIntoContext(context.Background(), "test-user")
	ctxWithInvalidUser := context.WithValue(context.Background(), "user", 123)

	invalidCreds := &credentials.Credentials{
		UserName: "",
		Password: "password",
		Type:     model.CREDENTIALS,
	}

	validCreds := &credentials.Credentials{
		UserName: "username",
		Password: "password",
		Type:     model.CREDENTIALS,
	}

	invalidCredsMeta := meta.Meta("")
	validCredsdMeta := meta.Meta("valid_attach")

	var tc = []struct {
		name         string
		ctx          context.Context
		creds        *credentials.Credentials
		credsMeta    meta.Meta
		encryptFunc  func([]byte) ([]byte, error)
		encryptCalls int
		saveFunc     func(ctx context.Context, creds *dto.Credentials) error
		saveCalls    int
		err          error
	}{
		{
			name:         "happy_path",
			ctx:          ctxWithUser,
			creds:        validCreds,
			credsMeta:    validCredsdMeta,
			encryptFunc:  func([]byte) ([]byte, error) { return []byte("encrypted"), nil },
			encryptCalls: 2,
			saveFunc:     func(ctx context.Context, creds *dto.Credentials) error { return nil },
			saveCalls:    1,
			err:          nil,
		},
		{
			name:      "invalid_user",
			ctx:       ctxWithInvalidUser,
			creds:     validCreds,
			credsMeta: validCredsdMeta,
			err:       ErrSomethingWentWrong,
		},
		{
			name:      "invalid_creds",
			ctx:       ctxWithUser,
			creds:     invalidCreds,
			credsMeta: validCredsdMeta,
			err:       ErrValidation,
		},
		{
			name:      "invalid_creds_meta",
			ctx:       ctxWithUser,
			creds:     validCreds,
			credsMeta: invalidCredsMeta,
			err:       ErrValidation,
		},
		{
			name:      "empty_user",
			ctx:       ctxWithoutUser,
			creds:     validCreds,
			credsMeta: validCredsdMeta,
			err:       ErrValidation,
		},
		{
			name:         "encrypt_error",
			ctx:          ctxWithUser,
			creds:        validCreds,
			credsMeta:    validCredsdMeta,
			encryptFunc:  func([]byte) ([]byte, error) { return nil, errors.New("test-error") },
			encryptCalls: 1,
			err:          ErrSomethingWentWrong,
		},
		{
			name:         "card_save_error",
			ctx:          ctxWithUser,
			creds:        validCreds,
			credsMeta:    validCredsdMeta,
			encryptFunc:  func(bytes []byte) ([]byte, error) { return []byte("encrypted"), nil },
			encryptCalls: 2,
			saveFunc:     func(ctx context.Context, creds *dto.Credentials) error { return errors.New("test-error") },
			saveCalls:    1,
			err:          ErrSomethingWentWrong,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			storage := &credentialsStorageMock{SaveCredentialsFunc: tt.saveFunc}

			encryptor := &encryptorMock{EncryptFunc: tt.encryptFunc}

			svc := &keepService{creds: storage, encryptor: encryptor}

			err := svc.KeepCredentials(tt.ctx, tt.creds, tt.credsMeta)
			require.ErrorIs(t, err, tt.err)
			require.Equal(t, tt.saveCalls, len(storage.calls.SaveCredentials))
			require.Equal(t, tt.encryptCalls, len(encryptor.calls.Encrypt))
		})
	}
}
