package order

import (
	"ddd-impl/config"
	"ddd-impl/internal/domains/customer"
	"ddd-impl/internal/domains/order"
	"ddd-impl/internal/domains/product"
)

type UsecaseOrder struct {
	repoOrder   order.MysqlOrderRepository
	repoProduct product.MysqlRepository
	cfg         *config.Config
}

func NewUsecaseOrder(repo order.MysqlOrderRepository) order.Usecase {
	return &UsecaseOrder{
		repoOrder: repo,
	}
}

func (u *UsecaseOrder) Create(p []*product.Product, c customer.Customer) (*order.Order, error) {
	for _, product := range p {
		return nil, product.GetPrice(product)
	}
	odr, err := order.NewOrder(p, c)
	odr.SetAdminFee(u.cfg.Additional.AdminFee)
	od, err := u.repoOrder.Save(odr)
	if err != nil {
		return nil, err
	}
	return od, nil
}

func (u *UsecaseOrder) UpdateOrder(status string, orderID int, promo, brand string) (*order.Order, error) {
	// Mendapatkan pesanan dari repository berdasarkan ID
	order, err := u.orderRepository.GetByID(orderID, promo, brand)
	if err != nil {
		return nil, err
	}

	// Mengatur status pesanan menggunakan setter
	order.Status = status

	// Simpan perubahan status ke dalam repository
	err = s.orderRepository.Update(order)
	if err != nil {
		return nil, err
	}

	return order, nil

}
