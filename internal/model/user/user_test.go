package user

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	var tc = []struct {
		name string
		att  User
		err  error
	}{
		{"empty user", User(""), ErrEmptyUser},
		{"valid", User("user"), nil},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.att.Validate()
			require.ErrorIs(t, tt.err, err)
		})
	}
}
