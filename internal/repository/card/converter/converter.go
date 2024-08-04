package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/card/dto"
)

func FromService(user user.User, meta meta.Meta) *dto.GetCard {
	return &dto.GetCard{
		User: string(user),
		Meta: string(meta),
	}
}

func ToService(c *dto.Card) *card.HashedCard {
	return &card.HashedCard{
		Owner:      c.Owner,
		NumberHash: c.NumberHash,
		NumberSalt: c.NumberSalt,
		CVVHash:    c.CVVHash,
		CVVSalt:    c.CVVSalt,
		Type:       model.CARD,
	}
}
