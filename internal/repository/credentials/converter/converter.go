package converter

import (
	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/types"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/credentials/model"
)

func RequestFromService(user user.User, meta meta.Meta) *model.GetCredentials {
	return &model.GetCredentials{
		User: string(user),
		Meta: string(meta),
	}
}

func CredentialsToService(creds *model.Credentials) *credentials.HashedCredentials {
	return &credentials.HashedCredentials{
		UserNameHash: creds.UserNameHash,
		UserNameSalt: creds.UserNameSalt,
		PasswordHash: creds.PasswordHash,
		PasswordSalt: creds.PasswordSalt,
		Type:         types.CREDENTIALS,
	}
}
