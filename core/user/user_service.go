package core

import (
	"errors"
	"fmt"
)

type UserService interface {
	GetAllUsers() ([]User, error)
	Register(user User) error
	UpdateUser(id string, user User) error
	DeleteUser(id string) error
	LoginUser(email, password string) (User, error)
}

type userServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) GetAllUsers() ([]User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userServiceImpl) Register(user User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Phone == "" {
		return errors.New("name and email and password and phone are required")
	}
	if err := s.repo.Register(user); err != nil {
		return err
	}
	return nil
}
func (s *userServiceImpl) UpdateUser(id string, user User) error {
	fmt.Printf("Updating user with ID: %s\n", id)
	if id == "" {
		return errors.New("user ID is required")
	}
	if err := s.repo.Update(id, user); err != nil {
		return err
	}
	return nil
}

func (s *userServiceImpl) DeleteUser(id string) error {
	if id == "" {
		return errors.New("user ID is required")
	}
	if err := s.repo.DeleteUser(id); err != nil {
		return err
	}
	return nil
}
func (s *userServiceImpl) LoginUser(email, password string) (User, error) {
	if email == "" || password == "" {
		return User{}, errors.New("email and password are required")
	}

	user, err := s.repo.LoginUser(email, password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
