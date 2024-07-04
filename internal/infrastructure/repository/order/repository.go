package order

import (
	"database/sql"
	"ddd-impl/internal/domains/order"
)

type Repository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) order.MysqlOrderRepository {
	return &Repository{
		db: db,
	}
}

func (r Repository) Save(order *order.Order) (*order.Order, error) {
	query := ""

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	exec, err := stmt.Exec(order.Customer.Name, ....)
	if err != nil {
		return nil, err
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return nil, err
	}

	order.Id = int(id)

	return order, nil
}

func (r Repository) Update(order *order.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r Repository) FindOrderByID(orderID uint) (*order.Order, error) {
	//TODO implement me
	panic("implement me")
}
