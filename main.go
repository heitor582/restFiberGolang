package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"teste.com/todos/src/configuration"
	"teste.com/todos/src/routes"
)

func init() {
	configuration.InitDatabase()
}

func main() {
	var PORT string = ":" + os.Getenv("PORT")
	app := fiber.New()
	routes.SetupUserRoutes(app)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}))
	routes.SetupTodoRoutes(app)
	fmt.Println(PORT)
	app.Listen(PORT)
}
