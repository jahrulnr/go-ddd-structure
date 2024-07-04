package order

import (
	"ddd-impl/internal/domains/customer"
	"ddd-impl/internal/domains/product"
	"time"
)

// Order aggregate
type Order struct {
	Id              int
	Customer        customer.Customer
	TotalPrice      int64              //
	TotalItem       int                //
	AdminFee        int64              //
	TransactionDate time.Time          //
	Items           []*product.Product //terdiri dari item dan quantity
	Status          string
}

// domain role
func NewOrder(p []*product.Product, c customer.Customer) (*Order, error) {

	// calcilate price
	var totalPrice int64
	for _, item := range p {
		totalPrice += item.Price
	}

	o := &Order{
		Id:              0,
		Customer:        c,
		AdminFee:        0,
		TotalPrice:      totalPrice,
		Items:           p,
		TransactionDate: time.Now(),
	}

	// set status pending
	return o, nil
}

func (o *Order) SetAdminFee(v int64) {
	// abil dari configrable
	o.AdminFee = v
}
