package converter

import (
	"bytes"
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/attachment/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRequestFromService(t *testing.T) {
	user := user.User("user")
	meta := meta.Meta("meta")

	req := FromService(user, meta)

	require.Equal(t, string(user), req.User)
	require.Equal(t, string(meta), req.Meta)
}

func TestAttachmentToService(t *testing.T) {
	attach := &dto.Attachment{
		User:    "user",
		Meta:    "meta",
		Name:    "name",
		Content: &bytes.Buffer{},
	}

	attach.Content.Write([]byte("content"))

	serviceAttach := ToService(attach)

	require.Equal(t, attach.Name, serviceAttach.Name)
	require.Equal(t, attach.Content.Bytes(), serviceAttach.Content)
	require.Equal(t, model.ATTACHMENT, serviceAttach.Type)
}
