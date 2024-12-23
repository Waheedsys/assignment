package handler

import (
	"github.com/Waheedsys/entities/entities"
)

type UserService interface {
	GetUsers() ([]entities.Users, error)
	GetUsersByName(name string) (entities.Users, error)
	AddUsers(user *entities.Users) error
	DeleteUsers(name string) error
	UpdateUsers(name string, updateUser *entities.Users) error
}
