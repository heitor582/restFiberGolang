package routes

import (
	"github.com/gofiber/fiber/v2"
	"teste.com/todos/src/controllers"
)

// SetupRoutes initialzize routes of app
func SetupTodoRoutes(app *fiber.App) {
	route := app.Group("/todos")
	route.Get("/", controllers.GetAllTodos)
	route.Get("/:id", controllers.GetTodo)
	route.Post("/", controllers.NewTodo)
	route.Delete("/:id", controllers.DeleteTodo)
}
