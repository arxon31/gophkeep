package card

import "errors"

var (
	ErrEmptyCardNumber = errors.New("card number cannot be empty")
	ErrEmptyOwner      = errors.New("owner cannot be empty")
	ErrEmptyCVV        = errors.New("cvv cannot be empty")
	ErrInvalidType     = errors.New("invalid type")
	ErrInvalidCVV      = errors.New("invalid cvv")
)
