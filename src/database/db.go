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

	connected := false

	for !connected {
		DB, err = gorm.Open(mysql.Open("admin:admin@tcp(db:3306)/meetup"), &gorm.Config{})

		if err != nil {
			// panic("Could not connect to database!")
			time.Sleep(1000)
		} else {
			connected = true
			break
		}

	}

}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
