package converter

import (
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/types"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"
)

func CardToService(c *gophkeep.SaveCardRequest) (*card.Card, meta.Meta) {
	return &card.Card{
		Owner:  c.GetOwner(),
		Number: c.GetNumber(),
		CVV:    c.GetCvv(),
		Type:   types.CARD,
	}, meta.Meta(c.Meta.GetMeta())
}

func CardToProto(c *card.Card) *gophkeep.GetCardResponse {
	return &gophkeep.GetCardResponse{
		Owner:      c.Owner,
		CardNumber: c.Number,
		Cvv:        c.CVV,
	}
}
