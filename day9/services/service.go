package services

import (
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
		return entities.Users{}, err
	}
	return user, nil
}

func (s *UserService) AddUsers(user *entities.Users) error {
	return s.store.AddUsers(user)
}

func (s *UserService) DeleteUsers(name string) error {
	return s.store.DeleteUsers(name)
}

func (s *UserService) UpdateUsers(name string, updateUser *entities.Users) error {
	return s.store.UpdateUsers(name, updateUser)
}
