package converter

import (
	"github.com/arxon31/yapr-proto/pkg/gophkeep"

	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/types"
)

func CredentialsToService(creds *gophkeep.SaveCredentialsRequest) (*credentials.Credentials, meta.Meta) {
	return &credentials.Credentials{
		UserName: creds.GetUsername(),
		Password: creds.GetPassword(),
		Type:     types.CREDENTIALS,
	}, meta.Meta(creds.Meta.GetMeta())
}

func CredentialsToProto(c *credentials.Credentials) *gophkeep.GetCredentialsResponse {
	return &gophkeep.GetCredentialsResponse{
		Username: c.UserName,
		Password: c.Password,
	}
}
