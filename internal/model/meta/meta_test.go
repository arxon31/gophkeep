package meta

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMeta_Validate(t *testing.T) {
	var tc = []struct {
		name string
		att  Meta
		err  error
	}{
		{"empty meta", Meta(""), ErrEmptyMeta},
		{"valid", Meta("meta"), nil},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.att.Validate()
			require.ErrorIs(t, tt.err, err)
		})
	}
}
