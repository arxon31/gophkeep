package keep

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	cardmodel "github.com/arxon31/gophkeep/internal/repository/card/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	. "github.com/arxon31/gophkeep/pkg/logger"
)

func (ks *keepService) KeepCard(ctx context.Context, card *card.Card, cardMeta meta.Meta) error {
	err := card.Validate()
	if err != nil {
		Logger.Error("card validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	err = cardMeta.Validate()
	if err != nil {
		Logger.Error("card meta validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	u, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		Logger.Error("extracting user from context", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	err = user.User(u).Validate()
	if err != nil {
		Logger.Error("user validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	numberHash, numberSalt, err := ks.hasher.Hash([]byte(card.Number))
	if err != nil {
		Logger.Error("card number hashing", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	cvvHash, cvvSalt, err := ks.hasher.Hash([]byte(fmt.Sprintf("%d", card.CVV)))
	if err != nil {
		Logger.Error("card cvv hashing", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	dbCard := &cardmodel.Card{
		User:       u,
		Meta:       string(cardMeta),
		Owner:      u,
		NumberHash: numberHash,
		NumberSalt: numberSalt,
		CVVHash:    cvvHash,
		CVVSalt:    cvvSalt,
	}

	err = ks.card.SaveCard(ctx, dbCard)
	if err != nil {
		Logger.Error("card saving", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	return nil
}
