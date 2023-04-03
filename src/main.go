package main

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/configs"
	"AgroXpert-Backend/src/database"
	"AgroXpert-Backend/src/routes"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":5000")

	configs.InitEnv()
	database.InitMongoConnection()
}
