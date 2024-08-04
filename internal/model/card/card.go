package card

import "github.com/arxon31/gophkeep/internal/model/types"

type Card struct {
	Owner  string
	Number string
	CVV    int64
	Type   int
}

type HashedCard struct {
	Owner      string
	NumberHash []byte
	NumberSalt []byte
	CVVHash    []byte
	CVVSalt    []byte
	Type       int
}

func (b Card) Validate() error {
	if b.Owner == "" {
		return ErrEmptyOwner
	}

	if b.Number == "" {
		return ErrEmptyCardNumber
	}

	if b.CVV == 0 {
		return ErrEmptyCVV
	}

	if b.CVV < 0 {
		return ErrInvalidCVV
	}

	if b.Type != types.CARD {
		return ErrInvalidType
	}

	return nil
}
