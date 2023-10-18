package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email" gorm:"unique"`
	Password     []byte  `json:"-"`
	IsAmbassador bool    `json:"-"`
	Revenue      float64 `json:"revenue,omitempty" gorm:"-"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return err
}
