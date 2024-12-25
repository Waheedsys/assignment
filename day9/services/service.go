package services

import (
	"database/sql"
	"fmt"
	"github.com/Waheedsys/entities/entities"
)

type UserService struct {
	store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) GetUsers() ([]entities.Users, error) {
	users, err := s.store.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) GetUsersByName(name string) (entities.Users, error) {
	user, err := s.store.GetUsersByName(name)
	if err != nil {
		if err == sql.ErrNoRows {

			return entities.Users{}, fmt.Errorf("user with name '%s' not found", name)
		}
		return entities.Users{}, err
	}

	return user, nil
}

func (s *UserService) AddUsers(user *entities.Users) error {
	existingUser, err := s.store.GetUsersByName(user.UserName)
	if err == nil && existingUser.UserName != "" {

		return fmt.Errorf("user with username '%s' already exists", user.UserName)
	}

	return s.store.AddUsers(user)
}

func (s *UserService) DeleteUsers(name string) error {
	existingUser, err := s.store.GetUsersByName(name)
	if err != nil || existingUser.UserName == "" {
		return fmt.Errorf("user with username '%s' does not exist", name)
	}

	return s.store.DeleteUsers(name)
}

func (s *UserService) UpdateUsers(name string, updateUser *entities.Users) error {
	existingUser, err := s.store.GetUsersByName(name)
	if err != nil || existingUser.UserName == "" {

		return fmt.Errorf("user with username '%s' does not exist", name)
	}

	return s.store.UpdateUsers(name, updateUser)
}
