package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCardToService(t *testing.T) {
	c := &gophkeep.SaveCardRequest{
		Meta:   &gophkeep.Meta{Meta: "meta"},
		Number: "123",
		Owner:  "owner",
		Cvv:    123,
	}

	card, meta := CardToService(c)

	require.Equal(t, c.Number, card.Number)
	require.Equal(t, c.Owner, card.Owner)
	require.Equal(t, c.Cvv, card.CVV)
	require.Equal(t, model.CARD, card.Type)

	require.Equal(t, "meta", string(meta))
}

func TestCardToProto(t *testing.T) {
	c := &card.Card{
		Owner:  "owner",
		Number: "123",
		CVV:    123,
		Type:   model.CARD,
	}

	protoCard := CardToProto(c)

	require.Equal(t, c.Owner, protoCard.Owner)
	require.Equal(t, c.Number, protoCard.CardNumber)
	require.Equal(t, c.CVV, protoCard.Cvv)
}
