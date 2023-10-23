package main

import (
	"context"
	"store/src/database"
	"store/src/models"

	"github.com/go-redis/redis/v8"
)

func main() {
	database.Connect()
	database.SetupRedis()

	ctx := context.Background()

	var users []models.User

	database.DB.Find(&users, models.User{
		IsAmbassador: true,
	})

	for _, user := range users {
		ambassador := models.Ambassador(user)
		ambassador.CalculateRevenue(database.DB)

		database.Cache.ZAdd(ctx, "rankings", &redis.Z{
			Score:  *ambassador.Revenue,
			Member: user.Name(),
		})
	}

}
