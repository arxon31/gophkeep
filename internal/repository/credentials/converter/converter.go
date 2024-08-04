package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/credentials/dto"
)

func FromService(user user.User, meta meta.Meta) *dto.GetCredentials {
	return &dto.GetCredentials{
		User: string(user),
		Meta: string(meta),
	}
}

func ToService(creds *dto.Credentials) *credentials.HashedCredentials {
	return &credentials.HashedCredentials{
		UserNameHash: creds.UserNameHash,
		UserNameSalt: creds.UserNameSalt,
		PasswordHash: creds.PasswordHash,
		PasswordSalt: creds.PasswordSalt,
		Type:         model.CREDENTIALS,
	}
}
