package order

type MysqlOrderRepository interface {
	Save(order *Order) (*Order, error)
	Update(order *Order) error
	FindOrderByID(orderID uint) (*Order, error)
}
