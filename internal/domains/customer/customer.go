package customer

type Customer struct {
	Id      string
	Name    string
	Email   string
	Address Address
}

type Address struct {
	Title      string
	Street     string
	City       string
	PostalCode string
}
