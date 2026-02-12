package core

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	Email      string    `gorm:"column:email" json:"email"`
	Name       string    `gorm:"column:name" json:"name"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	Address    string    `gorm:"column:address" json:"address"`
	Province   string    `gorm:"column:province" json:"province"`
	District   string    `gorm:"column:district" json:"district"`
	Image      string    `gorm:"column:image" json:"image"`
	IsActive   bool      `gorm:"column:isActive" json:"isActive"`
	IsVerified bool      `gorm:"column:isVerified" json:"isVerified"`
	Password   string    `gorm:"column:password" json:"password"`
	Role       string    `gorm:"column:role" json:"role"`
	CreatedAt  time.Time `gorm:"column:createdAt" json:"createdAt"`
	Village    string    `gorm:"column:village" json:"village"`
}

// TableName overrides the default table name to use the existing "User" table
func (User) TableName() string {
	return "User"
}
