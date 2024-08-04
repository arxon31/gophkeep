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

func ToService(c *dto.Card) *card.CryptedCard {
	return &card.CryptedCard{
		Owner:           c.Owner,
		EncryptedNumber: c.EncpryptedNumber,
		EncryptedCVV:    c.EncrpytedCVV,
		Type:            model.CARD,
	}
}
