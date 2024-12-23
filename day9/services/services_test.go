package services

import (
	"github.com/Waheedsys/entities/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetUsers(t *testing.T) {
	svc := NewUserService(mockStore{})

	resp, err := svc.GetUsers()
	user := []entities.Users{{UserName: "alice", UserAge: 23, Phone_number: "23542469786", Email: "waheed@12.com"},
		{UserName: "bob", UserAge: 33, Phone_number: "23500469786", Email: "bob@122.com"}}
	if resp[0] != user[0] {
		t.Errorf("expected %v, but got %v", user, resp)
	}
	if err != nil {
		t.Errorf("error while getting user")
	}

}
func Test_GetUsersByName(t *testing.T) {
	svc := NewUserService(mockStore{})
	resp, err := svc.GetUsersByName("alice")
	if err != nil {
		t.Errorf("error to get user")
	}
	mockuser := entities.Users{UserName: "alice", UserAge: 23, Phone_number: "23542469786", Email: "waheed@12.com"}

	if resp != mockuser {
		t.Errorf("expected %v, but got %v", mockuser, resp)
	}
	resp, err = svc.GetUsersByName("bob")
	if err == nil {
		assert.Equal(t, entities.Users{}, resp)
	}
	if err != nil {
		assert.Equal(t, entities.Users{}, resp)
	}
}

func Test_AddUsers(t *testing.T) {
	svc := NewUserService(mockStore{})
	mockuser := entities.Users{UserName: "alice", UserAge: 23, Phone_number: "23542469786", Email: "waheed@12.com"}
	err := svc.AddUsers(&mockuser)
	if err != nil {
		t.Errorf("error while calling adduser")
	}
}

func Test_DeleteUsers(t *testing.T) {
	svc := NewUserService(mockStore{})
	err := svc.DeleteUsers("alice")
	if err != nil {
		t.Errorf("error while deleting")
	}
}

func Test_UpdateUsers(t *testing.T) {
	svc := NewUserService(mockStore{})
	mockuser := entities.Users{UserName: "alice", UserAge: 23, Phone_number: "23542469786", Email: "waheed@12.com"}
	err := svc.UpdateUsers("alice", &mockuser)
	if err != nil {
		t.Errorf("error while deleting")
	}
}

type mockStore struct {
}

func (m mockStore) GetUsers() ([]entities.Users, error) {
	user := []entities.Users{{UserName: "alice", UserAge: 23, Phone_number: "23542469786", Email: "waheed@12.com"},
		{UserName: "bob", UserAge: 33, Phone_number: "23500469786", Email: "bob@122.com"}}
	return user, nil
}

func (m mockStore) GetUsersByName(name string) (entities.Users, error) {
	if name == "alice" {
		users := entities.Users{UserName: "alice", UserAge: 23, Phone_number: "23542469786", Email: "waheed@12.com"}
		return users, nil
	}

	if name == "bob" {
		return entities.Users{}, nil
	}

	return entities.Users{}, nil
}

func (m mockStore) AddUsers(user *entities.Users) error {
	return nil
}

func (m mockStore) DeleteUsers(name string) error {
	return nil
}

func (m mockStore) UpdateUsers(name string, updateUser *entities.Users) error {
	return nil
}
