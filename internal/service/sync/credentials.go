package sync

import (
	"context"
	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/credentials/converter"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	. "github.com/arxon31/gophkeep/pkg/logger"
	"log/slog"
)

func (ss *syncService) SyncCredentials(ctx context.Context, req *meta.Meta) (resp *credentials.Credentials, err error) {
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

	credentialsDB := converter.RequestFromService(user.User(u), *req)

	credsFromDB, err := ss.creds.GetCredentials(ctx, credentialsDB)
	if err != nil {
		Logger.Error("get attachment", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	hashedCreds := converter.CredentialsToService(credsFromDB)

	username, err := ss.unhasher.Unhash(hashedCreds.UserNameHash, hashedCreds.UserNameSalt)
	if err != nil {
		Logger.Error("username unhashing", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	password, err := ss.unhasher.Unhash(hashedCreds.PasswordHash, hashedCreds.PasswordSalt)
	if err != nil {
		Logger.Error("password unhashing", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	return &credentials.Credentials{
		UserName: username,
		Password: password,
		Type:     hashedCreds.Type,
	}, nil
}
