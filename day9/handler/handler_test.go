package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Waheedsys/entities/entities"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserService struct{}

// GetUser.
func TestUserHandler_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockUserService(ctrl)
	handler := NewUserHandler(mock)

	// Setup mock expectations for GetUsers()
	mock.EXPECT().GetUsers().Return([]entities.Users{
		{UserName: "alice", UserAge: 23, PhoneNumber: "23542469786", Email: "alice@example.com"},
	}, nil).Times(1) // Expect this call to be made once

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.GetUsers).Methods(http.MethodGet)

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var users []entities.Users
	if err := json.NewDecoder(rec.Body).Decode(&users); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(users) != 1 || users[0].UserName != "alice" {
		t.Errorf("expected user 'alice', got %v", users)
	}
}

// getbyname.
func TestUserHandler_GetUserByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockUserService(ctrl)
	handler := NewUserHandler(mock)

	req := httptest.NewRequest(http.MethodGet, "/user/alice", http.NoBody)
	rec := httptest.NewRecorder()
	r := mux.NewRouter()

	expectedUser := entities.Users{
		UserName:    "alice",
		UserAge:     23,
		PhoneNumber: "8090892381",
		Email:       "alice@gmail.com",
	}
	mock.EXPECT().GetUsersByName("alice").Return(expectedUser, nil)

	r.HandleFunc("/user/{name}", handler.GetUserByName).Methods(http.MethodGet)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var user entities.Users
	if err := json.NewDecoder(rec.Body).Decode(&user); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if user.UserName != "alice" {
		t.Errorf("expected user 'alice', got %v", user)
	}
}

// adduser.
func TestUserHandler_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockUserService(ctrl)
	handler := NewUserHandler(mock)

	newUser := entities.Users{
		UserName:    "alice",
		UserAge:     223,
		PhoneNumber: "8090892381",
		Email:       "alice@gmail.com",
	}

	body, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("failed to marshal new user: %v", err)
	}

	mock.EXPECT().AddUsers(&newUser).Return(nil)
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/user", handler.AddUser).Methods(http.MethodPost)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var addedUser entities.Users
	if err := json.NewDecoder(rec.Body).Decode(&addedUser); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if addedUser.UserName != newUser.UserName {
		t.Errorf("expected user '%s', got %v", newUser.UserName, addedUser)
	}
}

// Test: UpdateUser.
func TestUserHandler_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockUserService(ctrl)
	handler := NewUserHandler(mock)

	updatedUser := entities.Users{
		UserName:    "waheed",
		UserAge:     26,
		PhoneNumber: "987654321",
		Email:       "waheed@newexample.com",
	}

	body, err := json.Marshal(updatedUser)
	if err != nil {
		t.Fatalf("failed to marshal updated user: %v", err)
	}
	mock.EXPECT().UpdateUsers("waheed", &updatedUser).Return(nil)
	req := httptest.NewRequest(http.MethodPut, "/user/waheed", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", handler.UpdateUser).Methods(http.MethodPut)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var updated entities.Users
	if err := json.NewDecoder(rec.Body).Decode(&updated); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if updated.UserAge != 26 {
		t.Errorf("expected updated age to be 26, got %d", updated.UserAge)
	}
}

// Test: DeleteUser.
func TestUserHandler_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockUserService(ctrl)
	handler := NewUserHandler(mock)

	mock.EXPECT().DeleteUsers("waheed").Return(nil).Times(1)
	req := httptest.NewRequest(http.MethodDelete, "/user/waheed", http.NoBody)
	rec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", handler.DeleteUser).Methods(http.MethodDelete)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}

	// test notfound.
	mock.EXPECT().DeleteUsers("notfound").Return(fmt.Errorf("user not found")).Times(1)
	req = httptest.NewRequest(http.MethodDelete, "/user/notfound", nil)
	rec = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/user/{name}", handler.DeleteUser).Methods(http.MethodDelete)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

//
//func Test_UserHandler(t *testing.T) {
//
//	mockSvc := &mockUserService{}
//	handler := NewUserHandler(mockSvc)
//
//	req := httptest.NewRequest(http.MethodGet, "/users", http.NoBody)
//	rec := httptest.NewRecorder()
//	r := mux.NewRouter()
//	r.HandleFunc("/users", handler.GetUsers).Methods(http.MethodGet)
//	r.ServeHTTP(rec, req)
//
//	if rec.Code != http.StatusOK {
//		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
//	}
//
//	var users []entities.Users
//	if err := json.NewDecoder(rec.Body).Decode(&users); err != nil {
//		t.Fatalf("failed to decode response body: %v", err)
//	}
//
//	if len(users) != 1 || users[0].UserName != "waheed" {
//		t.Errorf("expected user 'waheed', got %v", users)
//	}
//
//	// getbyname
//	req = httptest.NewRequest(http.MethodGet, "/user/waheed", http.NoBody)
//	rec = httptest.NewRecorder()
//	r = mux.NewRouter()
//	r.HandleFunc("/user/{name}", handler.GetUserByName).Methods(http.MethodGet)
//	r.ServeHTTP(rec, req)
//
//	if rec.Code != http.StatusOK {
//		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
//	}
//
//	var user entities.Users
//	if err := json.NewDecoder(rec.Body).Decode(&user); err != nil {
//		t.Fatalf("failed to decode response body: %v", err)
//	}
//
//	if user.UserName != "waheed" {
//		t.Errorf("expected user 'waheed', got %v", user)
//	}
//
//	// Test: GetUserByName notfound
//	req = httptest.NewRequest(http.MethodGet, "/user/notfound", http.NoBody)
//	rec = httptest.NewRecorder()
//	r = mux.NewRouter()
//	r.HandleFunc("/user/{name}", handler.GetUserByName).Methods(http.MethodGet)
//	r.ServeHTTP(rec, req)
//
//	if rec.Code != http.StatusInternalServerError {
//		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
//	}
//
//	// Test: AddUser
//	newUser := entities.Users{
//		UserName:    "waheed",
//		UserAge:     25,
//		PhoneNumber: "123456789",
//		Email:       "waheed@example.com",
//	}
//
//	body, err := json.Marshal(newUser)
//	if err != nil {
//		t.Fatalf("failed to marshal new user: %v", err)
//	}
//
//	req = httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
//	rec = httptest.NewRecorder()
//	r = mux.NewRouter()
//	r.HandleFunc("/user", handler.AddUser).Methods(http.MethodPost)
//	r.ServeHTTP(rec, req)
//
//	if rec.Code != http.StatusOK {
//		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
//	}
//
//	var addedUser entities.Users
//	if err := json.NewDecoder(rec.Body).Decode(&addedUser); err != nil {
//		t.Fatalf("failed to decode response body: %v", err)
//	}
//
//	if addedUser.UserName != newUser.UserName {
//		t.Errorf("expected user '%s', got %v", newUser.UserName, addedUser)
//	}
//
//	// Test: UpdateUser
//	updatedUser := entities.Users{
//		UserName:    "waheed",
//		UserAge:     26,
//		PhoneNumber: "987654321",
//		Email:       "waheed@newexample.com",
//	}
//
//	body, err = json.Marshal(updatedUser)
//	if err != nil {
//		t.Fatalf("failed to marshal updated user: %v", err)
//	}
//
//	req = httptest.NewRequest(http.MethodPut, "/user/waheed", bytes.NewReader(body))
//	rec = httptest.NewRecorder()
//	r = mux.NewRouter()
//	r.HandleFunc("/user/{name}", handler.UpdateUser).Methods(http.MethodPut)
//	r.ServeHTTP(rec, req)
//
//	if rec.Code != http.StatusOK {
//		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
//	}
//
//	var updated entities.Users
//	if err := json.NewDecoder(rec.Body).Decode(&updated); err != nil {
//		t.Fatalf("failed to decode response body: %v", err)
//	}
//
//	if updated.UserAge != 26 {
//		t.Errorf("expected updated age to be 26, got %d", updated.UserAge)
//	}
//
//	// Test: DeleteUser
//	req = httptest.NewRequest(http.MethodDelete, "/user/waheed", http.NoBody)
//	rec = httptest.NewRecorder()
//	r = mux.NewRouter()
//	r.HandleFunc("/user/{name}", handler.DeleteUser).Methods(http.MethodDelete)
//	r.ServeHTTP(rec, req)
//
//	if rec.Code != http.StatusNoContent {
//		t.Errorf("expected status %d, got %d", http.StatusNoContent, rec.Code)
//	}
//
//	// Test: DeleteUser
//	req = httptest.NewRequest(http.MethodDelete, "/user/notfound", nil)
//	rec = httptest.NewRecorder()
//	r = mux.NewRouter()
//	r.HandleFunc("/user/{name}", handler.DeleteUser).Methods(http.MethodDelete)
//	r.ServeHTTP(rec, req)
//
//	if rec.Code != http.StatusInternalServerError {
//		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
//	}
//}

func (m *mockUserService) GetUsers() ([]entities.Users, error) {
	return []entities.Users{
		{
			UserName:    "alice",
			UserAge:     25,
			PhoneNumber: "123456789",
			Email:       "waheed@example.com",
		},
	}, nil
}

func (m *mockUserService) GetUsersByName(name string) (entities.Users, error) {
	if name == "alice" {
		return entities.Users{
			UserName:    "alice",
			UserAge:     25,
			PhoneNumber: "123456789",
			Email:       "waheed@example.com",
		}, nil
	}
	return entities.Users{}, fmt.Errorf("user not found")
}

func (m *mockUserService) AddUsers(user *entities.Users) error {
	if user.UserName == "alice" {
		return nil
	}
	return fmt.Errorf("unable to add user")
}

func (m *mockUserService) UpdateUsers(name string, updateUser *entities.Users) error {
	if name == "alice" {
		return nil
	}
	return fmt.Errorf("user not found")
}

func (m *mockUserService) DeleteUsers(name string) error {
	if name == "alice" {
		return nil
	}
	return fmt.Errorf("user not found")
}
