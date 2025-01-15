package converter

import (
	"github.com/arxon31/gophkeep/internal/model"
	"github.com/arxon31/gophkeep/internal/model/meta"
	"github.com/arxon31/gophkeep/internal/model/user"
	"github.com/arxon31/gophkeep/internal/repository/card/dto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromService(t *testing.T) {

	user := user.User("user")
	meta := meta.Meta("meta")

	req := FromService(user, meta)

	require.Equal(t, string(user), req.User)
	require.Equal(t, string(meta), req.Meta)
}

func TestToService(t *testing.T) {
	c := &dto.Card{
		User:             "user",
		Meta:             "meta",
		Owner:            "owner",
		EncpryptedNumber: []byte("number"),
		EncrpytedCVV:     []byte("cvv"),
	}

	svcCard := ToService(c)

	require.Equal(t, c.Owner, svcCard.Owner)
	require.Equal(t, c.EncpryptedNumber, svcCard.EncryptedNumber)
	require.Equal(t, c.EncrpytedCVV, svcCard.EncryptedCVV)

	require.Equal(t, model.CARD, svcCard.Type)
}
