package ctxfuncs

import (
	"context"
	"errors"
)

var (
	ErrUnknownUser = errors.New("unknown user type")
	ErrUnknownHash = errors.New("unknown hash type")
)

const (
	userKey        = "user"
	sessionHashKey = "session_hash"
)

func SetUserIntoContext(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUserFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(userKey)
	if val == nil {
		return "", nil
	}

	user, ok := val.(string)
	if !ok {
		return "", ErrUnknownUser
	}

	return user, nil
}

func GetSessionHashFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(sessionHashKey)
	if val == nil {
		return "", nil
	}

	sessionHash, ok := val.(string)
	if !ok {
		return "", ErrUnknownHash
	}

	return sessionHash, nil
}
