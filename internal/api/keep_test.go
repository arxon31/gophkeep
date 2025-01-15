package api

import (
	"context"
	"errors"
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/credentials"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServer_Keep(t *testing.T) {

	var srv *server = NewServer(nil, nil, nil)

	errMessage := "test-error"

	t.Run("card", func(t *testing.T) {
		var tc = []struct {
			name     string
			keepFunc func(ctx context.Context, card *card.Card, meta meta.Meta) error
			card     *gophkeep.SaveCardRequest
			status   *gophkeep.SaveStatus
			err      error
		}{
			{
				name: "happy_path",
				keepFunc: func(ctx context.Context, card *card.Card, meta meta.Meta) error {
					return nil
				},
				card: &gophkeep.SaveCardRequest{
					Owner:  "owner",
					Number: "123456",
					Cvv:    "123",
				},
				status: &gophkeep.SaveStatus{Success: true},
				err:    nil,
			},
			{
				name: "keep_error",
				keepFunc: func(ctx context.Context, card *card.Card, meta meta.Meta) error {
					return errors.New(errMessage)
				},
				card: &gophkeep.SaveCardRequest{
					Owner:  "owner",
					Number: "123456",
					Cvv:    "123",
				},
				status: &gophkeep.SaveStatus{Success: false, Error: &errMessage},
				err:    nil,
			},
		}
		for _, tt := range tc {
			t.Run(tt.name, func(t *testing.T) {
				service := &keepServiceMock{KeepCardFunc: tt.keepFunc}

				srv.keep = service

				status, err := srv.SaveCard(context.Background(), tt.card)
				if err != nil {
					require.ErrorIs(t, tt.err, err)
					require.Equal(t, status.Error, &errMessage)
				}
				require.Equal(t, tt.status, status)
				require.Equal(t, 1, len(service.calls.KeepCard))
			})
		}
	})

	t.Run("credentials", func(t *testing.T) {
		var tc = []struct {
			name     string
			keepFunc func(ctx context.Context, creds *credentials.Credentials, meta meta.Meta) error
			creds    *gophkeep.SaveCredentialsRequest
			status   *gophkeep.SaveStatus
			err      error
		}{
			{
				name: "happy_path",
				keepFunc: func(ctx context.Context, creds *credentials.Credentials, meta meta.Meta) error {
					return nil
				},
				creds: &gophkeep.SaveCredentialsRequest{
					Username: "username",
					Password: "password",
				},
				status: &gophkeep.SaveStatus{Success: true},
				err:    nil,
			},
			{
				name: "keep_error",
				keepFunc: func(ctx context.Context, creds *credentials.Credentials, meta meta.Meta) error {
					return errors.New(errMessage)
				},
				creds: &gophkeep.SaveCredentialsRequest{
					Username: "username",
					Password: "password",
				},
				status: &gophkeep.SaveStatus{Success: false, Error: &errMessage},
				err:    nil,
			},
		}
		for _, tt := range tc {
			t.Run(tt.name, func(t *testing.T) {
				service := &keepServiceMock{KeepCredentialsFunc: tt.keepFunc}

				srv.keep = service

				status, err := srv.SaveCredentials(context.Background(), tt.creds)
				if err != nil {
					require.ErrorIs(t, tt.err, err)
					require.Equal(t, status.Error, &errMessage)
				}
				require.Equal(t, tt.status, status)
				require.Equal(t, 1, len(service.calls.KeepCredentials))
			})
		}
	})
}
