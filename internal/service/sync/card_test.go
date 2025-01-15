package sync

import (
	"context"
	"errors"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/repository/card/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSyncService_SyncCard(t *testing.T) {
	ctxWithoutUser := context.Background()
	ctxWithUser := ctxfuncs.SetUserIntoContext(context.Background(), "test-user")
	ctxWithInvalidUser := context.WithValue(context.Background(), "user", 123)

	invalidCardMeta := meta.Meta("")
	validCardMeta := meta.Meta("valid_card")

	owner := "owner"
	number := "123123"
	cvv := "123"

	var tc = []struct {
		name         string
		ctx          context.Context
		meta         meta.Meta
		getFunc      func(ctx context.Context, card *dto.GetCard) (*dto.Card, error)
		getCalls     int
		decryptFunc  func([]byte) ([]byte, error)
		decryptCalls int
		err          error
	}{
		{
			name: "happy_path",
			ctx:  ctxWithUser,
			meta: validCardMeta,
			getFunc: func(ctx context.Context, card *dto.GetCard) (*dto.Card, error) {
				return &dto.Card{Owner: owner, EncpryptedNumber: []byte(number), EncrpytedCVV: []byte(cvv)}, nil
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
			meta: invalidCardMeta,
			err:  ErrValidation,
		},
		{
			name: "empty_user",
			ctx:  ctxWithoutUser,
			meta: validCardMeta,
			err:  ErrValidation,
		},
		{
			name: "get_card_error",
			ctx:  ctxWithUser,
			meta: validCardMeta,
			getFunc: func(ctx context.Context, card *dto.GetCard) (*dto.Card, error) {
				return nil, errors.New("some error")
			},
			getCalls: 1,
			err:      ErrSomethingWentWrong,
		},
		{
			name: "decrypt_error",
			ctx:  ctxWithUser,
			meta: validCardMeta,
			getFunc: func(ctx context.Context, card *dto.GetCard) (*dto.Card, error) {
				return &dto.Card{Owner: owner, EncpryptedNumber: []byte(number), EncrpytedCVV: []byte(cvv)}, nil
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
			provider := &cardProviderMock{GetCardFunc: tt.getFunc}
			decrypt := &decryptorMock{DecryptFunc: tt.decryptFunc}

			ss := &syncService{
				card:      provider,
				decryptor: decrypt,
			}

			card, err := ss.SyncCard(tt.ctx, &tt.meta)
			if err == nil {
				require.NotNil(t, card)

				require.Equal(t, card.Owner, owner)
				require.Equal(t, card.Number, number)
				require.Equal(t, card.CVV, cvv)
			}

			require.ErrorIs(t, err, tt.err)

			require.Equal(t, tt.getCalls, len(provider.calls.GetCard))
			require.Equal(t, tt.decryptCalls, len(decrypt.calls.Decrypt))
		})
	}
}
