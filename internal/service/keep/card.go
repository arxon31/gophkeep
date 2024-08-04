package keep

import (
	"context"
	"log/slog"

	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	cardmodel "github.com/arxon31/gophkeep/internal/repository/card/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	. "github.com/arxon31/gophkeep/pkg/logger"
)

func (ks *keepService) KeepCard(ctx context.Context, card *card.Card, cardMeta meta.Meta) error {
	u, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		Logger.Error("extracting user from context", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	err = card.Validate()
	if err != nil {
		Logger.Error("card validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	err = cardMeta.Validate()
	if err != nil {
		Logger.Error("card meta validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	err = user.User(u).Validate()
	if err != nil {
		Logger.Error("user validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	numberHash, err := ks.encryptor.Encrypt([]byte(card.Number))
	if err != nil {
		Logger.Error("card number encrypting", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	cvvHash, err := ks.encryptor.Encrypt([]byte(card.CVV))
	if err != nil {
		Logger.Error("card cvv encrypting", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	dbCard := &cardmodel.Card{
		User:             u,
		Meta:             string(cardMeta),
		Owner:            u,
		EncpryptedNumber: numberHash,
		EncrpytedCVV:     cvvHash,
	}

	err = ks.card.SaveCard(ctx, dbCard)
	if err != nil {
		Logger.Error("card saving", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	return nil
}
