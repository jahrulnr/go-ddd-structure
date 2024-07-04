package user

type UseCase interface {
	Register(user *User) (User, error)
}
