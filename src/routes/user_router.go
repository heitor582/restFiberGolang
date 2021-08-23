package routes

import (
	"github.com/gofiber/fiber/v2"
	"teste.com/todos/src/controllers"
)

// SetupRoutes initialzize routes of app
func SetupUserRoutes(app *fiber.App) {
	route := app.Group("/user")
	route.Post("/login", controllers.Login)
	route.Post("/register", controllers.RegisterUser)
}
