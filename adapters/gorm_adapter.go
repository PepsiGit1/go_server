package adapters

import (
	core "go_server/core/user"

	"gorm.io/gorm"
)

// Secondary adapter
type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) GetAll() ([]core.User, error) {
	var users []core.User
	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
