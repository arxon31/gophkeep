package credentials

import (
	"github.com/arxon31/gophkeep/internal/model"
)

type Credentials struct {
	UserName string
	Password string
	Type     int
}

type HashedCredentials struct {
	UserNameHash []byte
	UserNameSalt []byte
	PasswordHash []byte
	PasswordSalt []byte
	Type         int
}

func (c Credentials) Validate() error {
	if c.UserName == "" {
		return ErrEmptyUserName
	}

	if c.Password == "" {
		return ErrEmptyPassword
	}

	if c.Type != model.CREDENTIALS {
		return ErrInvalidType
	}

	return nil
}
