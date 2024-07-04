package product

type Product struct {
	Id       int
	Name     string
	Price    int64
	Discount int
}

func (p *Product) GetPrice(product *Product) error {
	if product.Price == 0 {
		return ErrPrice
	}
	return nil
}
