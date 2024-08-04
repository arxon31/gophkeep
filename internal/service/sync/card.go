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
	"strconv"
)

func (ss *syncService) SyncCard(ctx context.Context, req *meta.Meta) (resp *card.Card, err error) {
	err = req.Validate()
	if err != nil {
		Logger.Error("attachment meta validation", slog.String("error", err.Error()))
		return nil, ErrValidation
	}

	u, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		Logger.Error("extracting user from context", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	err = user.User(u).Validate()
	if err != nil {
		Logger.Error("user validation", slog.String("error", err.Error()))
		return nil, ErrValidation
	}

	cardDB := converter.RequestFromService(user.User(u), *req)

	cardFromDB, err := ss.card.GetCard(ctx, cardDB)
	if err != nil {
		Logger.Error("get attachment", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	hashedCard := converter.CardToService(cardFromDB)

	cvvString, err := ss.unhasher.Unhash(hashedCard.CVVHash, hashedCard.CVVSalt)
	if err != nil {
		Logger.Error("card cvv unhashing", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	cvv, err := strconv.ParseInt(cvvString, 10, 32)
	if err != nil {
		Logger.Error("card cvv parsing", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	number, err := ss.unhasher.Unhash(hashedCard.NumberHash, hashedCard.NumberSalt)
	if err != nil {
		Logger.Error("card number unhashing", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	return &card.Card{
		Owner:  hashedCard.Owner,
		Number: number,
		CVV:    cvv,
		Type:   hashedCard.Type,
	}, nil

}
