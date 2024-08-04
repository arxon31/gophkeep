package sync

import (
	"context"
	"github.com/arxon31/gophkeep/internal/model/attachment"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/attachment/converter"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	. "github.com/arxon31/gophkeep/pkg/logger"
	"log/slog"
)

func (ss *syncService) SyncAttachment(ctx context.Context, req *meta.Meta) (*attachment.Attachment, error) {
	err := req.Validate()
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

	attachDB := converter.RequestFromService(user.User(u), *req)

	attach, err := ss.attachments.GetAttachment(ctx, attachDB)
	if err != nil {
		Logger.Error("get attachment", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	return converter.AttachmentToService(attach), nil
}
