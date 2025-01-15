package attachment

import "errors"

var (
	ErrEmptyName    = errors.New("name cannot be empty")
	ErrEmptyContent = errors.New("content cannot be empty")
	ErrInvalidType  = errors.New("invalid type")
)
