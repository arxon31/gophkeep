package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/attachment"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/attachment/dto"
)

func FromService(user user.User, meta meta.Meta) *dto.GetAttachment {
	return &dto.GetAttachment{
		User: string(user),
		Meta: string(meta),
	}
}

func ToService(attach *dto.Attachment) *attachment.Attachment {
	return &attachment.Attachment{
		Name:    attach.Name,
		Content: attach.Content.Bytes(),
		Type:    model.ATTACHMENT,
	}
}
