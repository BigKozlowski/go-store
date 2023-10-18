package models

import (
	"store/src/database"

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

type Admin User

type Ambassador User

func (ambassador *Ambassador) CalculateRevenue() {
	var orders []Order

	database.DB.Preload("OrderItems").Find(&orders, &Order{
		UserId:   ambassador.Id,
		Complete: true,
	})

	var revenue float64 = 0

	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AdminRevenue
		}
	}

	ambassador.Revenue = revenue
}

func (admin *Admin) CalculateRevenue() {

}
