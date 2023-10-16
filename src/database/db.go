package database

import (
	"meetup/src/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(mysql.Open("admin:admin@tcp(db:3306)/meetup"), &gorm.Config{})

		if err != nil {
			time.Sleep(time.Second)
		} else {
			return
		}
	}
	panic("Could not connect to database!")
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
