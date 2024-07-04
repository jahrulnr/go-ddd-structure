package order

import "errors"

var (
	ErrName  = errors.New("name cannot be empty")
	ErrEmail = errors.New("email cannot be empty")
)
