package entities

import (
	"errors"
	"regexp"
)

type Users struct {
	UserName     string `json:"user_name"`
	UserAge      int    `json:"user_age"`
	Phone_number string `json:"phone_Number"`
	Email        string `json:"email"`
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
	isPhoneValid := phone.MatchString(u.Phone_number)
	if !isPhoneValid {
		return errors.New("Invalid phone number")
	}

	return nil

}
