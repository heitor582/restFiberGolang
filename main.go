package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"teste.com/todos/src/configuration"
	"teste.com/todos/src/routes"
)

func init() {
	configuration.InitDatabase()
}

func main() {
	var PORT string = ":" + os.Getenv("PORT")
	app := fiber.New()
	app.Group("/api/v1")
	routes.SetupTodoRoutes(app)
	fmt.Println(PORT)
	app.Listen(PORT)
}
