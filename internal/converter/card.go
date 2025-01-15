package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/yapr-proto/pkg/gophkeep"

	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
)

func CardToService(c *gophkeep.SaveCardRequest) (*card.Card, meta.Meta) {
	return &card.Card{
		Owner:  c.GetOwner(),
		Number: c.GetNumber(),
		CVV:    c.GetCvv(),
		Type:   model.CARD,
	}, meta.Meta(c.Meta.GetMeta())
}

func CardToProto(c *card.Card) *gophkeep.GetCardResponse {
	return &gophkeep.GetCardResponse{
		Owner:      c.Owner,
		CardNumber: c.Number,
		Cvv:        c.CVV,
	}
}
