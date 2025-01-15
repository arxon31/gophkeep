package credentials

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCredentials_Validate(t *testing.T) {
	var tc = []struct {
		name string
		att  *Credentials
		err  error
	}{
		{"empty username", &Credentials{}, ErrEmptyUserName},
		{"empty password", &Credentials{UserName: "username"}, ErrEmptyPassword},
		{"invalid type", &Credentials{UserName: "username", Password: "password"}, ErrInvalidType},
		{"valid", &Credentials{UserName: "username", Password: "password", Type: model.CREDENTIALS}, nil},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.att.Validate()
			require.ErrorIs(t, tt.err, err)
		})
	}
}
