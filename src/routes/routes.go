package routes

import (
	"store/src/controllers"
	"store/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	admin := api.Group("admin")
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("/user", controllers.User)
	adminAuthenticated.Post("/logout", controllers.Logout)
	adminAuthenticated.Put("users/info", controllers.UpdateInfo)
	adminAuthenticated.Put("users/password", controllers.UpdatePassword)
	adminAuthenticated.Get("ambassadors", controllers.Ambassadors)
}
