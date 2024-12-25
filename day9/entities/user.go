package entities

import (
	"errors"
	"regexp"
)

var ErrInvalidPhoneNumber = errors.New("invalid phone number")

type Users struct {
	UserName    string `json:"user_name"`
	UserAge     int    `json:"user_age"`
	PhoneNumber string `json:"phone_Number"`
	Email       string `json:"email"`
}

func (u *Users) Validate(users Users) error {

	emailRe := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	email := regexp.MustCompile(emailRe)
	isEmailValid := email.MatchString(u.Email)
	if !isEmailValid {
		return errors.New("Invalid email id")
	}

	phoneRe := `^\+?[0-9]{10,15}$`
	phone := regexp.MustCompile(phoneRe)
	isPhoneValid := phone.MatchString(u.PhoneNumber)
	if !isPhoneValid {
		return ErrInvalidPhoneNumber
	}

	return nil

}
