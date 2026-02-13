package adapters

import (
	core "go_server/core/user"

	"gorm.io/gorm"
)

// Secondary adapter
type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) core.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) GetAll() ([]core.User, error) {
	var users []core.User
	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *GormUserRepository) Register(user core.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormUserRepository) Update(id string, user core.User) error {
	result := r.db.Model(&core.User{}).
		Where("id = ?", id).
		Updates(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormUserRepository) DeleteUser(id string) error {
	result := r.db.Delete(&core.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormUserRepository) LoginUser(email, password string) (core.User, error) {
	var user core.User
	result := r.db.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return core.User{}, result.Error
	}
	return user, nil
}
