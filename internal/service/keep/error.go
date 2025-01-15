package keep

import "errors"

var (
	ErrValidation         = errors.New("invalid request")
	ErrSomethingWentWrong = errors.New("something went wrong")
)
