package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Waheedsys/entities/entities"
	"github.com/Waheedsys/entities/services"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	user, err := h.UserService.GetUsersByName(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser entities.Users
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.UserService.AddUsers(&newUser); err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var updateUser entities.Users
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.UserService.UpdateUsers(name, &updateUser); err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateUser)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if err := h.UserService.DeleteUsers(name); err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
