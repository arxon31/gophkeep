package keep

import (
	"context"
	"log/slog"

	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	credsmodel "github.com/arxon31/gophkeep/internal/repository/credentials/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	. "github.com/arxon31/gophkeep/pkg/logger"
)

func (ks *keepService) KeepCredentials(ctx context.Context, creds *credentials.Credentials, credsMeta meta.Meta) error {
	u, err := ctxfuncs.GetUserFromContext(ctx)
	if err != nil {
		Logger.Error("extracting user from context", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	err = creds.Validate()
	if err != nil {
		Logger.Error("credentials validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	err = credsMeta.Validate()
	if err != nil {
		Logger.Error("credentials meta validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	err = user.User(u).Validate()
	if err != nil {
		Logger.Error("user validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	encryptedUserName, err := ks.encryptor.Encrypt([]byte(creds.UserName))
	if err != nil {
		Logger.Error("username encrypting", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	encryptedPassword, err := ks.encryptor.Encrypt([]byte(creds.Password))
	if err != nil {
		Logger.Error("password encrypting", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	dbCreds := &credsmodel.Credentials{
		User:              u,
		Meta:              string(credsMeta),
		EncryptedPassword: encryptedPassword,
		EncryptedUserName: encryptedUserName,
	}

	err = ks.creds.SaveCredentials(ctx, dbCreds)
	if err != nil {
		Logger.Error("credentials saving", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	return nil
}
