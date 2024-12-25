package stores

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Waheedsys/entities/entities"
)

type UsersList struct {
	db *sql.DB
}

func NewDetails(db *sql.DB) *UsersList {
	return &UsersList{db: db}
}

// getuser.
func (userStore *UsersList) GetUsers() ([]entities.Users, error) {
	rows, err := userStore.db.Query("SELECT UserName,UserAge,PhoneNumber,Email FROM User")
	if err != nil {
		log.Printf("Error :%v", err)
		return nil, err
	}

	var users []entities.Users

	for rows.Next() {
		var user entities.Users

		if err := rows.Scan(&user.UserName, &user.PhoneNumber, &user.UserAge, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

var ErrUserNotFound = errors.New("user not found")

// GetUsersByName retrieves a user from the store by their name.
func (userStore *UsersList) GetUsersByName(name string) (entities.Users, error) {
	var user entities.Users
	err := userStore.db.QueryRow("SELECT UserName, UserAge, PhoneNumber, Email FROM User WHERE Username = ?", name).
		Scan(&user.UserName, &user.UserAge, &user.PhoneNumber, &user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return entities.Users{}, fmt.Errorf("%w:user with username '%s' not found", ErrUserNotFound, name)
	}

	return user, nil
}

// GetUsersByName adds's a user to the store.
func (userStore *UsersList) AddUsers(user *entities.Users) error {
	log.Printf("Inserting user: %+v", user)
	_, err := userStore.db.Exec("INSERT INTO User (UserName, UserAge, PhoneNumber, Email) VALUES (?, ?, ?, ?)",
		user.UserName, user.UserAge, user.PhoneNumber, user.Email)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
	}

	return err
}

// DeleteUsers removes a user from the UsersList based on the provided username (name).
// It returns an error if the user cannot be deleted.
//
// Parameters:
//   - name: The username of the user to be deleted.
//
// Returns:
//   - error: Returns nil if the deletion is successful, otherwise returns an error.
func (userStore *UsersList) DeleteUsers(name string) error {
	_, err := userStore.db.Exec("DELETE FROM User WHERE UserName = ?", name)
	return err
}

// UpdateUsers updates the details of an existing user in the UsersList.
// It takes a username (name) and a pointer to an updated Users struct (updateUser),
// and returns an error if the update fails.
//
// Parameters:
//   - name: The username of the user to be updated.
//   - updateUser: The updated user details.
//
// Returns:
//   - error: Returns nil if the update is successful, otherwise returns an error.
func (userStore *UsersList) UpdateUsers(name string, updateUser *entities.Users) error {
	_, err := userStore.db.Exec("UPDATE User SET Email = ? WHERE UserName = ?", updateUser.Email, name)
	return err
}
