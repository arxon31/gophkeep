package credentials

import "errors"

var (
	ErrEmptyUserName = errors.New("username cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrInvalidType   = errors.New("invalid type")
)
