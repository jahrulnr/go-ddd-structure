package user

type User struct {
	Id       int
	Name     string
	Username string
	Email    string
}

func NewUser() *User {
	//
	return &User{}
}
