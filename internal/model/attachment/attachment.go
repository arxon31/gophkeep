package attachment

import (
	"github.com/arxon31/gophkeep/internal/model"
)

type Attachment struct {
	Name    string
	Content []byte
	Type    int
}

func (a *Attachment) Validate() error {
	if a.Name == "" {
		return ErrEmptyName
	}

	if a.Content == nil {
		return ErrEmptyContent
	}

	if a.Type != model.ATTACHMENT {
		return ErrInvalidType
	}

	return nil
}
