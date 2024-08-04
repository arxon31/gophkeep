package sync

import (
	"bytes"
	"context"
	"errors"
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/repository/attachment/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSyncService_SyncAttachment(t *testing.T) {
	ctxWithoutUser := context.Background()
	ctxWithUser := ctxfuncs.SetUserIntoContext(context.Background(), "test-user")
	ctxWithInvalidUser := context.WithValue(context.Background(), "user", 123)

	invalidAttachMeta := meta.Meta("")
	validAttachMeta := meta.Meta("valid_attach")

	name := "test-name"
	content := []byte("test-content")

	var tc = []struct {
		name    string
		ctx     context.Context
		meta    meta.Meta
		getFunc func(ctx context.Context, attachment *dto.GetAttachment) (*dto.Attachment, error)
		calls   int
		err     error
	}{
		{
			name: "happy_path",
			ctx:  ctxWithUser,
			meta: validAttachMeta,
			getFunc: func(ctx context.Context, attachment *dto.GetAttachment) (*dto.Attachment, error) {
				return &dto.Attachment{
					Name:    name,
					Content: bytes.NewBuffer(content),
				}, nil
			},
			calls: 1,
			err:   nil,
		},

		{
			name: "invalid_user",
			ctx:  ctxWithInvalidUser,
			err:  ErrSomethingWentWrong,
		},

		{
			name: "invalid_meta",
			ctx:  ctxWithUser,
			meta: invalidAttachMeta,
			err:  ErrValidation,
		},

		{
			name: "empty_user",
			ctx:  ctxWithoutUser,
			meta: validAttachMeta,
			err:  ErrValidation,
		},
		{
			name: "attachment_get_error",
			ctx:  ctxWithUser,
			meta: validAttachMeta,
			getFunc: func(ctx context.Context, attachment *dto.GetAttachment) (*dto.Attachment, error) {
				return nil, errors.New("test-error")
			},
			calls: 1,
			err:   ErrSomethingWentWrong,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			provider := &attachmentsProviderMock{
				GetAttachmentFunc: tt.getFunc,
			}

			svc := syncService{attachments: provider}

			attach, err := svc.SyncAttachment(tt.ctx, &tt.meta)
			if err == nil {
				require.NotNil(t, attach)
				require.Equal(t, attach.Type, model.ATTACHMENT)
				require.Equal(t, attach.Name, name)
				require.Equal(t, attach.Content, content)
			}
			require.ErrorIs(t, err, tt.err)
			require.Equal(t, tt.calls, len(provider.calls.GetAttachment))
		})
	}

}
