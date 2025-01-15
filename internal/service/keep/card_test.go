package keep

import (
	"context"
	"errors"
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/repository/card/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeepService_SaveCard(t *testing.T) {
	ctxWithoutUser := context.Background()
	ctxWithUser := ctxfuncs.SetUserIntoContext(context.Background(), "test-user")
	ctxWithInvalidUser := context.WithValue(context.Background(), "user", 123)

	invalidCard := &card.Card{
		Owner:  "",
		Number: "123",
		CVV:    "123",
		Type:   model.CARD,
	}

	validCard := &card.Card{
		Owner:  "owner",
		Number: "123",
		CVV:    "123",
		Type:   model.CARD,
	}

	invalidCardMeta := meta.Meta("")
	validCardMeta := meta.Meta("valid_attach")

	var tc = []struct {
		name         string
		ctx          context.Context
		card         *card.Card
		cardMeta     meta.Meta
		encryptFunc  func([]byte) ([]byte, error)
		encryptCalls int
		saveFunc     func(ctx context.Context, card *dto.Card) error
		saveCalls    int
		err          error
	}{
		{
			name:         "happy_path",
			ctx:          ctxWithUser,
			card:         validCard,
			cardMeta:     validCardMeta,
			encryptFunc:  func([]byte) ([]byte, error) { return []byte("encrypted"), nil },
			encryptCalls: 2,
			saveFunc:     func(ctx context.Context, card *dto.Card) error { return nil },
			saveCalls:    1,
			err:          nil,
		},
		{
			name:     "invalid_user",
			ctx:      ctxWithInvalidUser,
			card:     validCard,
			cardMeta: validCardMeta,
			err:      ErrSomethingWentWrong,
		},
		{
			name:     "invalid_card",
			ctx:      ctxWithUser,
			card:     invalidCard,
			cardMeta: validCardMeta,
			err:      ErrValidation,
		},
		{
			name:     "invalid_card_meta",
			ctx:      ctxWithUser,
			card:     validCard,
			cardMeta: invalidCardMeta,
			err:      ErrValidation,
		},
		{
			name:     "empty_user",
			ctx:      ctxWithoutUser,
			card:     validCard,
			cardMeta: validCardMeta,
			err:      ErrValidation,
		},
		{
			name:         "encrypt_error",
			ctx:          ctxWithUser,
			card:         validCard,
			cardMeta:     validCardMeta,
			encryptFunc:  func([]byte) ([]byte, error) { return nil, errors.New("test-error") },
			encryptCalls: 1,
			err:          ErrSomethingWentWrong,
		},
		{
			name:         "card_save_error",
			ctx:          ctxWithUser,
			card:         validCard,
			cardMeta:     validCardMeta,
			encryptFunc:  func(bytes []byte) ([]byte, error) { return []byte("encrypted"), nil },
			encryptCalls: 2,
			saveFunc:     func(ctx context.Context, card *dto.Card) error { return errors.New("test-error") },
			saveCalls:    1,
			err:          ErrSomethingWentWrong,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			storage := &cardStorageMock{SaveCardFunc: tt.saveFunc}

			encryptor := &encryptorMock{EncryptFunc: tt.encryptFunc}

			svc := &keepService{card: storage, encryptor: encryptor}

			err := svc.KeepCard(tt.ctx, tt.card, tt.cardMeta)
			require.ErrorIs(t, err, tt.err)
			require.Equal(t, tt.saveCalls, len(storage.calls.SaveCard))
			require.Equal(t, tt.encryptCalls, len(encryptor.calls.Encrypt))
		})
	}
}
