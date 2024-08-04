package keep

import (
	"bytes"
	"context"
	"log/slog"

	"github.com/arxon31/gophkeep/internal/model/attachment"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	attachmodel "github.com/arxon31/gophkeep/internal/repository/attachment/model"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	. "github.com/arxon31/gophkeep/pkg/logger"
)

func (ks *keepService) KeepAttachment(ctx context.Context, attach *attachment.Attachment, attachMeta meta.Meta) error {
	err := attach.Validate()
	if err != nil {
		Logger.Error("attach validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	err = attachMeta.Validate()
	if err != nil {
		Logger.Error("attach meta validation", slog.String("error", err.Error()))
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

	dbAttach := &attachmodel.Attachment{
		User:    u,
		Meta:    string(attachMeta),
		Name:    attach.Name,
		Content: bytes.NewBuffer(attach.Content),
	}

	err = ks.attachs.SaveAttachment(ctx, dbAttach)
	if err != nil {
		Logger.Error("attach saving", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	return nil
}
