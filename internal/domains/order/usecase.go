package order

import (
	"ddd-impl/internal/domains/customer"
	"ddd-impl/internal/domains/product"
)

type Usecase interface {
	Create([]*product.Product, customer.Customer) (*Order, error)
}
