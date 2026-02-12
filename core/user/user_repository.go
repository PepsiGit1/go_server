package core

// Secondary port
type OrderRepository interface {
	GetAll() ([]User, error)
}
