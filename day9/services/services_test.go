package services

import (
	"github.com/Waheedsys/entities/entities"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := NewMockUserStore(ctrl)

	svc := NewUserService(mock)
	user := []entities.Users{{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"},
		{UserName: "bob", UserAge: 33, PhoneNumber: "23500469786", Email: "bob@122.com"}}

	mock.EXPECT().GetUsers().Return(user, nil)

	resp, err := svc.GetUsers()

	if len(resp) != len(user) {
		t.Errorf("expected %v, but got %v", user, resp)
	}
	if resp[0] != user[0] {
		t.Errorf("expected %v, but got %v", user, resp)
	}
	if err != nil {
		t.Errorf("error while getting user")
	}
	for i, user1 := range user {

		if resp[i] != user1 {
			t.Errorf("expected %v, but got %v at index %d", user1, resp[i], i)
		}
	}

}

func Test_GetUsersByName(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := NewMockUserStore(ctrl)

	svc := NewUserService(mock)

	user := []entities.Users{{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"},
		{UserName: "bob", UserAge: 33, PhoneNumber: "23500469786", Email: "bob@122.com"}}

	mock.EXPECT().GetUsersByName("alice").Return(user[0], nil)

	resp, err := svc.GetUsersByName("alice")

	if resp != user[0] {
		t.Errorf("expected %v, but got %v", user, resp)
	}
	if err != nil {
		t.Errorf("error while getting GetUsersByName")
	}

}

func Test_AddUsers(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := NewMockUserStore(ctrl)

	svc := NewUserService(mock)

	user := []entities.Users{{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"},
		{UserName: "bob", UserAge: 33, PhoneNumber: "23500469786", Email: "bob@122.com"}}

	mock.EXPECT().GetUsersByName("bob").Return(entities.Users{}, nil)

	mock.EXPECT().AddUsers(&user[1]).Return(nil)

	err := svc.AddUsers(&user[1])
	if err != nil {
		t.Errorf("error while calling addusers")
	}

}

func Test_DeleteUsers(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := NewMockUserStore(ctrl)

	svc := NewUserService(mock)
	user := []entities.Users{{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"},
		{UserName: "bob", UserAge: 33, PhoneNumber: "23500469786", Email: "bob@122.com"}}

	mock.EXPECT().GetUsersByName("alice").Return(user[0], nil)
	mock.EXPECT().DeleteUsers("alice").Return(nil)

	err := svc.DeleteUsers("alice")
	if err != nil {
		t.Errorf("error while calling DeleteUsers")
	}

}

func Test_UpdateUsers(t *testing.T) {
	ctrl := gomock.NewController(t)

	mock := NewMockUserStore(ctrl)

	svc := NewUserService(mock)

	user := []entities.Users{{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"},
		{UserName: "bob", UserAge: 33, PhoneNumber: "23500469786", Email: "bob@122.com"}}

	mock.EXPECT().GetUsersByName("alice").Return(user[0], nil)
	mock.EXPECT().UpdateUsers("alice", &user[0]).Return(nil)

	err := svc.UpdateUsers("alice", &user[0])
	if err != nil {
		t.Errorf("error while calling updating")
	}

}

//	func Test_GetUsers(t *testing.T) {
//		svc := NewUserService(mockStore{})
//
//		resp, err := svc.GetUsers()
//		user := []entities.Users{{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"},
//			{UserName: "bob", UserAge: 33, PhoneNumber: "23500469786", Email: "bob@122.com"}}
//		if resp[0] != user[0] {
//			t.Errorf("expected %v, but got %v", user, resp)
//		}
//		if err != nil {
//			t.Errorf("error while getting user")
//		}
//
// }

//func Test_GetUsersByName(t *testing.T) {
//	svc := NewUserService(mockStore{})
//	resp, err := svc.GetUsersByName("alice")
//	if err != nil {
//		t.Errorf("error to get user")
//	}
//	mockuser := entities.Users{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"}
//
//	if resp != mockuser {
//		t.Errorf("expected %v, but got %v", mockuser, resp)
//	}
//	resp, err = svc.GetUsersByName("bob")
//	if err == nil {
//		assert.Equal(t, entities.Users{}, resp)
//	}
//	if err != nil {
//		assert.Equal(t, entities.Users{}, resp)
//	}
//}

// old test adduser
//func Test_AddUsers(t *testing.T) {
//	svc := NewUserService(mockStore{})
//	mockuser := entities.Users{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"}
//	err := svc.AddUsers(&mockuser)
//	if err != nil {
//		t.Errorf("error while calling adduser")
//	}
//}

//func Test_DeleteUsers(t *testing.T) {
//	svc := NewUserService(mockStore{})
//	err := svc.DeleteUsers("alice")
//	if err != nil {
//		t.Errorf("error while deleting")
//	}
//}

//func Test_UpdateUsers(t *testing.T) {
//	svc := NewUserService(mockStore{})
//	mockuser := entities.Users{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"}
//	err := svc.UpdateUsers("alice", &mockuser)
//	if err != nil {
//		t.Errorf("error while updating")
//	}
//}

type mockStore struct {
}

func (m mockStore) GetUsers() ([]entities.Users, error) {
	user := []entities.Users{{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"},
		{UserName: "bob", UserAge: 33, PhoneNumber: "23500469786", Email: "bob@122.com"}}
	return user, nil
}

func (m mockStore) GetUsersByName(name string) (entities.Users, error) {
	if name == "alice" {
		users := entities.Users{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "waheed@12.com"}
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
