package sync

import (
	"context"
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/card/converter"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	. "github.com/arxon31/gophkeep/pkg/logger"
	"log/slog"
)

func (ss *syncService) SyncCard(ctx context.Context, req *meta.Meta) (*card.Card, error) {
	u, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		Logger.Error("extracting user from context", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	err = req.Validate()
	if err != nil {
		Logger.Error("attachment meta validation", slog.String("error", err.Error()))
		return nil, ErrValidation
	}

	err = user.User(u).Validate()
	if err != nil {
		Logger.Error("user validation", slog.String("error", err.Error()))
		return nil, ErrValidation
	}

	cardDB := converter.FromService(user.User(u), *req)

	cardFromDB, err := ss.card.GetCard(ctx, cardDB)
	if err != nil {
		Logger.Error("get card", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	encryptedCard := converter.ToService(cardFromDB)

	cvv, err := ss.decryptor.Decrypt(encryptedCard.EncryptedCVV)
	if err != nil {
		Logger.Error("card cvv decrypting", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	number, err := ss.decryptor.Decrypt(encryptedCard.EncryptedNumber)
	if err != nil {
		Logger.Error("card number decrypting", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	return &card.Card{
		Owner:  encryptedCard.Owner,
		Number: string(number),
		CVV:    string(cvv),
		Type:   encryptedCard.Type,
	}, nil

}
