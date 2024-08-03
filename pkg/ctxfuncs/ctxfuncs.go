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
	val := ctx.Value("user")
	if val == nil {
		return "", nil
	}

	user, ok := val.(string)
	if !ok {
		return "", ErrUnknownUser
	}

	return user, nil
}
