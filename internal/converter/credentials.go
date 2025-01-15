package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"

	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
)

func CredentialsToService(creds *gophkeep.SaveCredentialsRequest) (*credentials.Credentials, meta.Meta) {
	return &credentials.Credentials{
		UserName: creds.GetUsername(),
		Password: creds.GetPassword(),
		Type:     model.CREDENTIALS,
	}, meta.Meta(creds.Meta.GetMeta())
}

func CredentialsToProto(c *credentials.Credentials) *gophkeep.GetCredentialsResponse {
	return &gophkeep.GetCredentialsResponse{
		Username: c.UserName,
		Password: c.Password,
	}
}
