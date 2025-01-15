package keep

import (
	"context"
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/attachment"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/repository/attachment/dto"
	"github.com/arxon31/gophkeep/pkg/ctxfuncs"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeepService_KeepAttachment(t *testing.T) {
	ctxWithoutUser := context.Background()
	ctxWithUser := ctxfuncs.SetUserIntoContext(context.Background(), "test-user")
	ctxWithInvalidUser := context.WithValue(context.Background(), "user", 123)

	invalidAttach := &attachment.Attachment{
		Name:    "invalid_attach",
		Content: nil,
		Type:    model.ATTACHMENT,
	}

	validAttach := &attachment.Attachment{
		Name:    "valid_attach",
		Content: []byte("content"),
		Type:    model.ATTACHMENT,
	}

	invalidAttachMeta := meta.Meta("")
	validAttachMeta := meta.Meta("valid_attach")

	var tc = []struct {
		name       string
		ctx        context.Context
		attach     *attachment.Attachment
		attachMeta meta.Meta
		saveFunc   func(ctx context.Context, attach *dto.Attachment) error
		calls      int
		err        error
	}{
		{
			name:       "happy_path",
			ctx:        ctxWithUser,
			attach:     validAttach,
			attachMeta: validAttachMeta,
			saveFunc:   func(ctx context.Context, attach *dto.Attachment) error { return nil },
			calls:      1,
			err:        nil,
		},
		{
			name:       "invalid_user",
			ctx:        ctxWithInvalidUser,
			attach:     validAttach,
			attachMeta: validAttachMeta,
			err:        ErrSomethingWentWrong,
		},
		{
			name:       "empty_user",
			ctx:        ctxWithoutUser,
			attach:     validAttach,
			attachMeta: validAttachMeta,
			err:        ErrValidation,
		},
		{
			name:       "invalid_attach",
			ctx:        ctxWithUser,
			attach:     invalidAttach,
			attachMeta: validAttachMeta,
			err:        ErrValidation,
		},
		{
			name:       "invalid_attach_meta",
			ctx:        ctxWithUser,
			attach:     validAttach,
			attachMeta: invalidAttachMeta,
			err:        ErrValidation,
		},
		{
			name:       "save_func_error",
			ctx:        ctxWithUser,
			attach:     validAttach,
			attachMeta: validAttachMeta,
			saveFunc:   func(ctx context.Context, attach *dto.Attachment) error { return ErrSomethingWentWrong },
			calls:      1,
			err:        ErrSomethingWentWrong,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			storage := &attachmentsStorageMock{
				SaveAttachmentFunc: tt.saveFunc,
			}
			ks := &keepService{attachs: storage}

			err := ks.KeepAttachment(tt.ctx, tt.attach, tt.attachMeta)
			require.ErrorIs(t, err, tt.err)
			require.Equal(t, tt.calls, len(storage.calls.SaveAttachment))
		})
	}

}
