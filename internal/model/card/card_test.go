package card

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCard_Validate(t *testing.T) {
	var tc = []struct {
		name string
		att  *Card
		err  error
	}{
		{"empty name", &Card{}, ErrEmptyOwner},
		{"empty number", &Card{Owner: "owner"}, ErrEmptyCardNumber},
		{"empty cvv", &Card{Owner: "owner", Number: "number"}, ErrEmptyCVV},
		{"invalid cvv", &Card{Owner: "owner", Number: "number", CVV: -1}, ErrInvalidCVV},
		{"invalid type", &Card{Owner: "owner", Number: "number", CVV: 1}, ErrInvalidType},
		{"valid", &Card{Owner: "owner", Number: "number", CVV: 1, Type: model.CARD}, nil},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.att.Validate()
			require.ErrorIs(t, tt.err, err)
		})
	}
}
