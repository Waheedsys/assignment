package services

import (
	"github.com/Waheedsys/entities/entities"
	"github.com/Waheedsys/entities/stores"
)

type Userservice interface {
	GetUsers() ([]entities.Users, error)
	GetUsersByName(name string) (entities.Users, error)
	AddUsers(user *entities.Users) error
	DeleteUsers(name string) error
	UpdateUsers(name string, updateUser *entities.Users) error
}
type UserService struct {
	store stores.UsersList
}

func NewUserService(store stores.UsersList) *UserService {
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
	//var userval entities.Users
	//
	//if err := userval.Validate(user); err != nil {
	//	return err
	//}

	return s.store.AddUsers(user)
}

func (s *UserService) DeleteUsers(name string) error {
	return s.store.DeleteUsers(name)
}

func (s *UserService) UpdateUsers(name string, updateUser *entities.Users) error {
	//var obj entities.Users
	//if err := obj.Validate(updateUser); err != nil {
	//	return err
	//}

	return s.store.UpdateUsers(name, updateUser)
}
