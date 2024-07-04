package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `gorm:"primaryKey" validate:"required"`
	Password string `gorm:"not null" validate:"required"`
}

/*
Hashing isn't enough to secure the password, it should be at least salted.
Since this is a simple example, we'll only use bcrypt for hashing the password.
Please don't trust this application with your real passwords.
For more information, see: https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html
*/

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
