package utils

import (
	"store/src/database"
	"store/src/models"
)

type Admin models.User

type Ambassador models.User

func (ambassador *Ambassador) CalculateRevenue() {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders, &models.Order{
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
