package converter

import (
	"github.com/arxon31/gophkeep/internal/model/attachment"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/types"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/attachment/model"
)

func RequestFromService(user user.User, meta meta.Meta) *model.GetAttachment {
	return &model.GetAttachment{
		User: string(user),
		Meta: string(meta),
	}
}

func AttachmentToService(attach *model.Attachment) *attachment.Attachment {
	return &attachment.Attachment{
		Name:    attach.Name,
		Content: attach.Content.Bytes(),
		Type:    types.ATTACHMENT,
	}
}
