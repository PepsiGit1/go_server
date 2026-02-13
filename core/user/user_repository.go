package core

// Secondary port
type UserRepository interface {
	GetAll() ([]User, error)
	Register(user User) error
	Update(id string, user User) error
	DeleteUser(id string) error
	LoginUser(email, password string) (User, error)
}
