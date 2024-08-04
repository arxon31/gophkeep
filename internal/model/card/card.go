package card

import (
	"github.com/arxon31/gophkeep/internal/model"
)

type Card struct {
	Owner  string
	Number string
	CVV    string
	Type   int
}

type CryptedCard struct {
	Owner           string
	EncryptedNumber []byte
	EncryptedCVV    []byte
	Type            int
}

func (b Card) Validate() error {
	if b.Owner == "" {
		return ErrEmptyOwner
	}

	if b.Number == "" {
		return ErrEmptyCardNumber
	}

	if b.CVV == "" {
		return ErrEmptyCVV
	}

	if b.Type != model.CARD {
		return ErrInvalidType
	}

	return nil
}
