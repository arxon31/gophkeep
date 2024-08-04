package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCredentialsToService(t *testing.T) {
	c := &gophkeep.SaveCredentialsRequest{
		Meta:     &gophkeep.Meta{Meta: "meta"},
		Username: "username",
		Password: "password",
	}

	creds, meta := CredentialsToService(c)

	require.Equal(t, c.Username, creds.UserName)
	require.Equal(t, c.Password, creds.Password)
	require.Equal(t, model.CREDENTIALS, creds.Type)

	require.Equal(t, "meta", string(meta))
}

func TestCredentialsToProto(t *testing.T) {
	c := &credentials.Credentials{
		UserName: "username",
		Password: "password",
		Type:     model.CREDENTIALS,
	}

	protoCreds := CredentialsToProto(c)

	require.Equal(t, c.UserName, protoCreds.Username)
	require.Equal(t, c.Password, protoCreds.Password)
}
