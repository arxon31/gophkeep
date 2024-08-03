package converter

import (
	"github.com/arxon31/gophkeep/internal/model/card"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/types"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/card/model"
)

func RequestFromService(user user.User, meta meta.Meta) *model.GetCard {
	return &model.GetCard{
		User: string(user),
		Meta: string(meta),
	}
}

func CardToService(c *model.Card) *card.HashedCard {
	return &card.HashedCard{
		Owner:      c.Owner,
		NumberHash: c.NumberHash,
		NumberSalt: c.NumberSalt,
		CVVHash:    c.CVVHash,
		CVVSalt:    c.CVVSalt,
		Type:       types.CARD,
	}
}
