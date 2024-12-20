package stores

import (
	"database/sql"
	"fmt"
	"github.com/Waheedsys/entities/entities"
	"log"
)

type UsersList struct {
	db *sql.DB
}

func NewDetails(db *sql.DB) *UsersList {
	return &UsersList{db: db}
}

// getuser
func (userStore *UsersList) GetUsers() ([]entities.Users, error) {

	rows, err := userStore.db.Query("SELECT UserName,UserAge,Phone_number,Email FROM User")
	if err != nil {
		log.Printf("Error :%v", err)
		return nil, err
	}
	var users []entities.Users
	for rows.Next() {
		var user entities.Users
		if err := rows.Scan(&user.UserName, &user.Phone_number, &user.UserAge, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUsersByName
func (userStore *UsersList) GetUsersByName(name string) (entities.Users, error) {
	var user entities.Users
	err := userStore.db.QueryRow("SELECT UserName, UserAge, Phone_number, Email FROM User WHERE Username = ?", name).
		Scan(&user.UserName, &user.UserAge, &user.Phone_number, &user.Email)
	if err == sql.ErrNoRows {
		return entities.Users{}, fmt.Errorf("user with username '%s' not found", name)
	}

	if err != nil {
		return entities.Users{}, fmt.Errorf("error querying user: %v", err)
	}

	return user, nil

}

// AddUsers
func (userStore *UsersList) AddUsers(user *entities.Users) error {
	log.Printf("Inserting user: %+v", user)
	_, err := userStore.db.Exec("INSERT INTO User (UserName, UserAge, Phone_number, Email) VALUES (?, ?, ?, ?)",
		user.UserName, user.UserAge, user.Phone_number, user.Email)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
	}
	return err
}

// DeleteUsers
func (userStore *UsersList) DeleteUsers(name string) error {
	_, err := userStore.db.Exec("DELETE FROM User WHERE UserName = ?", name)
	return err
}

// UpdateUsers
func (userStore *UsersList) UpdateUsers(name string, updateUser *entities.Users) error {
	_, err := userStore.db.Exec("UPDATE User SET Email = ? WHERE UserName = ?", updateUser.Email, name)
	return err
}
