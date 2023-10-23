package controllers

import (
	"context"
	"store/src/database"
	"store/src/models"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func Ambassadors(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}

type ScoreEntry struct {
	Name    string
	Revenue float64
}

func Rankings(c *fiber.Ctx) error {
	rankings, err := database.Cache.ZRevRangeByScoreWithScores(context.Background(), "rankings", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil {
		return err
	}

	var result []ScoreEntry

	for _, ranking := range rankings {
		result = append(result, ScoreEntry{
			Name:    ranking.Member.(string),
			Revenue: ranking.Score,
		})
	}

	return c.JSON(result)
}
