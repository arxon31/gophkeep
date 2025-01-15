package session

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type someInfo struct {
	ID   int
	Info string
}

func TestSessionService(t *testing.T) {
	svc := NewService()
	var sessID string

	t.Run("must_create", func(t *testing.T) {
		sessID = svc.Create(&someInfo{ID: 1, Info: "some info"})
		require.NotEmpty(t, sessID)
	})

	t.Run("must_return_info", func(t *testing.T) {
		info, ok := svc.Info(sessID)
		require.True(t, ok)
		require.IsType(t, &someInfo{}, info)

		i := info.(*someInfo)
		require.Equal(t, 1, i.ID)
		require.Equal(t, "some info", i.Info)
	})

	t.Run("must_delete", func(t *testing.T) {
		svc.Delete(sessID)
	})

	t.Run("must_return_nil", func(t *testing.T) {
		info, ok := svc.Info(sessID)
		require.False(t, ok)
		require.Nil(t, info)
	})

}
