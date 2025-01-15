package sync

import (
	"context"
	"errors"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/repository/credentials/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSyncService_SyncCredentials(t *testing.T) {
	ctxWithoutUser := context.Background()
	ctxWithUser := ctxfuncs.SetUserIntoContext(context.Background(), "test-user")
	ctxWithInvalidUser := context.WithValue(context.Background(), "user", 123)

	invalidCredsMeta := meta.Meta("")
	validCredsMeta := meta.Meta("valid_credentials")

	username := "username"
	password := "password"

	var tc = []struct {
		name         string
		ctx          context.Context
		meta         meta.Meta
		getFunc      func(ctx context.Context, card *dto.GetCredentials) (*dto.Credentials, error)
		getCalls     int
		decryptFunc  func([]byte) ([]byte, error)
		decryptCalls int
		err          error
	}{
		{
			name: "happy_path",
			ctx:  ctxWithUser,
			meta: validCredsMeta,
			getFunc: func(ctx context.Context, card *dto.GetCredentials) (*dto.Credentials, error) {
				return &dto.Credentials{EncryptedUserName: []byte(username), EncryptedPassword: []byte(password)}, nil
			},
			getCalls: 1,
			decryptFunc: func(bytes []byte) ([]byte, error) {
				return bytes, nil
			},
			decryptCalls: 2,
			err:          nil,
		},

		{
			name: "invalid_user",
			ctx:  ctxWithInvalidUser,
			err:  ErrSomethingWentWrong,
		},
		{
			name: "invalid_meta",
			ctx:  ctxWithUser,
			meta: invalidCredsMeta,
			err:  ErrValidation,
		},
		{
			name: "empty_user",
			ctx:  ctxWithoutUser,
			meta: validCredsMeta,
			err:  ErrValidation,
		},
		{
			name: "get_card_error",
			ctx:  ctxWithUser,
			meta: validCredsMeta,
			getFunc: func(ctx context.Context, card *dto.GetCredentials) (*dto.Credentials, error) {
				return nil, errors.New("some error")
			},
			getCalls: 1,
			err:      ErrSomethingWentWrong,
		},
		{
			name: "decrypt_error",
			ctx:  ctxWithUser,
			meta: validCredsMeta,
			getFunc: func(ctx context.Context, card *dto.GetCredentials) (*dto.Credentials, error) {
				return &dto.Credentials{EncryptedUserName: []byte(username), EncryptedPassword: []byte(password)}, nil
			},
			getCalls: 1,
			decryptFunc: func(bytes []byte) ([]byte, error) {
				return nil, errors.New("some error")
			},
			decryptCalls: 1,
			err:          ErrSomethingWentWrong,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			provider := &credentialsProviderMock{GetCredentialsFunc: tt.getFunc}
			decrypt := &decryptorMock{DecryptFunc: tt.decryptFunc}

			ss := &syncService{
				creds:     provider,
				decryptor: decrypt,
			}

			creds, err := ss.SyncCredentials(tt.ctx, &tt.meta)
			if err == nil {
				require.NotNil(t, creds)

				require.Equal(t, creds.UserName, username)
				require.Equal(t, creds.Password, password)
			}

			require.ErrorIs(t, err, tt.err)

			require.Equal(t, tt.getCalls, len(provider.calls.GetCredentials))
			require.Equal(t, tt.decryptCalls, len(decrypt.calls.Decrypt))
		})
	}
}
