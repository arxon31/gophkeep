package ctxfuncs

import (
	"context"
	"errors"
)

var ErrUnknownUser = errors.New("unknown user")

const userKey = "user"

func SetUserIntoContext(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUserFromContext(ctx context.Context) (string, error) {
	user, ok := ctx.Value("user").(string)
	if !ok {
		return "", ErrUnknownUser
	}
	return user, nil
}
