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

func (ss *syncService) SyncCredentials(ctx context.Context, req *meta.Meta) (*credentials.Credentials, error) {
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

	credentialsDB := converter.FromService(user.User(u), *req)

	credsFromDB, err := ss.creds.GetCredentials(ctx, credentialsDB)
	if err != nil {
		Logger.Error("get attachment", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	encryptedCreds := converter.ToService(credsFromDB)

	username, err := ss.decryptor.Decrypt(encryptedCreds.EncryptedUserName)
	if err != nil {
		Logger.Error("username decrypting", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	password, err := ss.decryptor.Decrypt(encryptedCreds.EncryptedPassword)
	if err != nil {
		Logger.Error("password decrypting", slog.String("error", err.Error()))
		return nil, ErrSomethingWentWrong
	}

	return &credentials.Credentials{
		UserName: string(username),
		Password: string(password),
		Type:     encryptedCreds.Type,
	}, nil
}
