package controllers

import (
	"context"
	"encoding/json"
	"sort"
	"store/src/database"
	"store/src/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Products(c *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	id, _ := strconv.Atoi(c.Params("id"))

	product.Id = uint(id)

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product
	product.Id = uint(id)

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var product models.Product
	product.Id = uint(id)

	database.DB.Delete(&product)

	return nil
}

func ProductsFrontend(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "products_frontend").Result()
	if err != nil {

		database.DB.Find(&products)

		bytes, err := json.Marshal(products)

		if err != nil {
			panic(err)
		}

		database.Cache.Set(ctx, "products_frontend", bytes, 30*time.Minute)
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	return c.JSON(products)
}

func ProductsBackend(c *fiber.Ctx) error {
	var products []models.Product
	var ctx = context.Background()

	result, err := database.Cache.Get(ctx, "products_backend").Result()
	if err != nil {

		database.DB.Find(&products)

		bytes, err := json.Marshal(products)

		if err != nil {
			panic(err)
		}

		database.Cache.Set(ctx, "products_backend", bytes, 30*time.Minute)
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	var foundProducts []models.Product

	if s := c.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), lower) || strings.Contains(strings.ToLower(product.Description), lower) {
				foundProducts = append(foundProducts, product)
			}
		}
	} else {
		foundProducts = products
	}

	if sortParam := c.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(foundProducts, func(i, j int) bool {
				return foundProducts[i].Price < foundProducts[j].Price
			})
		} else if sortLower == "dsc" {
			sort.Slice(foundProducts, func(i, j int) bool {
				return foundProducts[i].Price > foundProducts[j].Price
			})
		}
	}

	return c.JSON(foundProducts)
}
