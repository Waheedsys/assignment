package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Waheedsys/entities/entities"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserService struct{}

func Test_UserHandler(t *testing.T) {

	mockSvc := &mockUserService{}
	handler := NewUserHandler(mockSvc)

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

	if len(users) != 1 || users[0].UserName != "waheed" {
		t.Errorf("expected user 'waheed', got %v", users)
	}

	//getbynmae
	req = httptest.NewRequest(http.MethodGet, "/user/waheed", nil)
	rec = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/user/{name}", handler.GetUserByName).Methods(http.MethodGet)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var user entities.Users
	if err := json.NewDecoder(rec.Body).Decode(&user); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if user.UserName != "waheed" {
		t.Errorf("expected user 'waheed', got %v", user)
	}

	// Test: GetUserByName
	req = httptest.NewRequest(http.MethodGet, "/user/notfound", nil)
	rec = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/user/{name}", handler.GetUserByName).Methods(http.MethodGet)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}

	// Test: AddUser
	newUser := entities.Users{
		UserName:     "waheed",
		UserAge:      25,
		Phone_number: "123456789",
		Email:        "waheed@example.com",
	}

	body, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("failed to marshal new user: %v", err)
	}

	req = httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader(body))
	rec = httptest.NewRecorder()
	r = mux.NewRouter()
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

	// Test: UpdateUser
	updatedUser := entities.Users{
		UserName:     "waheed",
		UserAge:      26,
		Phone_number: "987654321",
		Email:        "waheed@newexample.com",
	}

	body, err = json.Marshal(updatedUser)
	if err != nil {
		t.Fatalf("failed to marshal updated user: %v", err)
	}

	req = httptest.NewRequest(http.MethodPut, "/user/waheed", bytes.NewReader(body))
	rec = httptest.NewRecorder()
	r = mux.NewRouter()
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

	// Test: DeleteUser
	req = httptest.NewRequest(http.MethodDelete, "/user/waheed", nil)
	rec = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/user/{name}", handler.DeleteUser).Methods(http.MethodDelete)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}

	// Test: DeleteUser
	req = httptest.NewRequest(http.MethodDelete, "/user/notfound", nil)
	rec = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/user/{name}", handler.DeleteUser).Methods(http.MethodDelete)
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rec.Code)
	}
}

func (m *mockUserService) GetUsers() ([]entities.Users, error) {
	return []entities.Users{
		{
			UserName:     "waheed",
			UserAge:      25,
			Phone_number: "123456789",
			Email:        "waheed@example.com",
		},
	}, nil
}

func (m *mockUserService) GetUsersByName(name string) (entities.Users, error) {
	if name == "waheed" {
		return entities.Users{
			UserName:     "waheed",
			UserAge:      25,
			Phone_number: "123456789",
			Email:        "waheed@example.com",
		}, nil
	}
	return entities.Users{}, fmt.Errorf("user not found")
}

func (m *mockUserService) AddUsers(user *entities.Users) error {
	if user.UserName == "waheed" {
		return nil
	}
	return fmt.Errorf("unable to add user")
}

func (m *mockUserService) UpdateUsers(name string, updateUser *entities.Users) error {
	if name == "waheed" {
		return nil
	}
	return fmt.Errorf("user not found")
}

func (m *mockUserService) DeleteUsers(name string) error {
	if name == "waheed" {
		return nil
	}
	return fmt.Errorf("user not found")
}
