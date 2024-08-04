package keep

import (
	"context"

	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	credsmodel "github.com/arxon31/gophkeep/internal/repository/credentials/model"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
)

func (ks *keepService) KeepCredentials(ctx context.Context, creds *credentials.Credentials, credsMeta meta.Meta) error {
	err := creds.Validate()
	if err != nil {
		return ErrValidation
	}

	err = credsMeta.Validate()
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

	usernameHash, usernameSalt, err := ks.hasher.Hash([]byte(creds.UserName))
	if err != nil {
		return ErrSomethingWentWrong
	}

	passwordHash, passwordSalt, err := ks.hasher.Hash([]byte(creds.Password))
	if err != nil {
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
		return ErrSomethingWentWrong
	}

	return nil
}
