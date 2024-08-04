package keep

import (
	"context"
	"log/slog"

	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	credsmodel "github.com/arxon31/gophkeep/internal/repository/credentials/model"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	. "github.com/arxon31/gophkeep/pkg/logger"
)

func (ks *keepService) KeepCredentials(ctx context.Context, creds *credentials.Credentials, credsMeta meta.Meta) error {
	err := creds.Validate()
	if err != nil {
		Logger.Error("credentials validation", slog.String("error", err.Error()))
		return ErrValidation
	}

	err = credsMeta.Validate()
	if err != nil {
		Logger.Error("credentials meta validation", slog.String("error", err.Error()))
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

	usernameHash, usernameSalt, err := ks.hasher.Hash([]byte(creds.UserName))
	if err != nil {
		Logger.Error("username hashing", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	passwordHash, passwordSalt, err := ks.hasher.Hash([]byte(creds.Password))
	if err != nil {
		Logger.Error("password hashing", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	dbCreds := &credsmodel.Credentials{
		User:         u,
		Meta:         string(credsMeta),
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
		UserNameHash: usernameHash,
		UserNameSalt: usernameSalt,
	}

	err = ks.creds.SaveCredentials(ctx, dbCreds)
	if err != nil {
		Logger.Error("credentials saving", slog.String("error", err.Error()))
		return ErrSomethingWentWrong
	}

	return nil
}
