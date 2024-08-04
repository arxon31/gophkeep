package keep

import (
	"context"
	"fmt"
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	cardmodel "github.com/arxon31/gophkeep/internal/repository/card/model"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
)

func (ks *keepService) SaveCard(ctx context.Context, card *card.Card, cardMeta meta.Meta) error {
	err := card.Validate()
	if err != nil {
		return ErrValidation
	}

	err = cardMeta.Validate()
	if err != nil {
		return ErrValidation
	}

	u, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		return ErrSomethingWentWrong
	}

	err = user.User(u).Validate()
	if err != nil {
		return ErrValidation
	}

	numberHash, numberSalt, err := ks.hasher.Hash([]byte(card.Number))
	if err != nil {
		return ErrSomethingWentWrong
	}

	cvvHash, cvvSalt, err := ks.hasher.Hash([]byte(fmt.Sprintf("%d", card.CVV)))
	if err != nil {
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
		return ErrSomethingWentWrong
	}

	return nil
}
