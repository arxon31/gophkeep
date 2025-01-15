package attachment

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttachment_Validate(t *testing.T) {
	var tc = []struct {
		name string
		att  *Attachment
		err  error
	}{
		{"empty name", &Attachment{}, ErrEmptyName},
		{"empty content", &Attachment{Name: "name"}, ErrEmptyContent},
		{"invalid type", &Attachment{Name: "name", Content: []byte("content")}, ErrInvalidType},
		{"valid", &Attachment{Name: "name", Content: []byte("content"), Type: model.ATTACHMENT}, nil},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.att.Validate()
			require.ErrorIs(t, tt.err, err)
		})
	}
}
