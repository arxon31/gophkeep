package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/credentials/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromService(t *testing.T) {

	user := user.User("user")
	meta := meta.Meta("meta")

	req := FromService(user, meta)

	require.Equal(t, string(user), req.User)
	require.Equal(t, string(meta), req.Meta)
}

func TestToService(t *testing.T) {
	creds := &dto.Credentials{
		User:              "user",
		Meta:              "meta",
		EncryptedUserName: []byte("username"),
		EncryptedPassword: []byte("password"),
	}

	svcCreds := ToService(creds)

	require.Equal(t, creds.EncryptedUserName, svcCreds.EncryptedUserName)
	require.Equal(t, creds.EncryptedPassword, svcCreds.EncryptedPassword)

	require.Equal(t, model.CREDENTIALS, svcCreds.Type)

}
