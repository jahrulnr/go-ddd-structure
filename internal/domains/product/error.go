package product

import "errors"

var (
	ErrPrice = errors.New("Price cannot be zero")
)
