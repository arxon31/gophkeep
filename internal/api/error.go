package api

import "errors"

var (
	ErrSessionNotFound    = errors.New("session not found")
	ErrSomethingWentWrong = errors.New("something went wrong")
)
