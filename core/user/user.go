package core

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	Phone      string     `json:"phone"`
	Address    string     `json:"address"`
	Province   string     `json:"province"`
	District   string     `json:"district"`
	Image      *string    `json:"image,omitempty"` // nullable
	IsActive   bool       `json:"isActive"`
	IsVerified bool       `json:"isVerified"`
	Password   string     `json:"password"`
	Role       string     `json:"role"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"` // optional
	Village    *string    `json:"village,omitempty"`   // optional
}

// TableName overrides the default table name to use the existing "User" table
func (User) TableName() string {
	return "User"
}
