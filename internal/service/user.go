package service

import (
	"errors"

	"github.com/yourusername/go_server/internal/models"
)

func (s *Service) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *Service) GetUserByID(id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.GetUserByID(id)
}

func (s *Service) CreateUser(user *models.User) error {
	// Add business logic validation here
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Name == "" {
		return errors.New("name is required")
	}

	return s.repo.CreateUser(user)
}

func (s *Service) UpdateUser(user *models.User) error {
	// Add business logic validation here
	if user.ID == "" {
		return errors.New("invalid user ID")
	}

	return s.repo.UpdateUser(user)
}

func (s *Service) DeleteUser(id string) error {
	if id == "" {
		return errors.New("invalid user ID")
	}
	return s.repo.DeleteUser(id)
}
