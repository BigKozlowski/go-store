package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	Ismeetup  bool
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return err
}
