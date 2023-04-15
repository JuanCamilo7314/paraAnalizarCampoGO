package main

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/configs"
	"AgroXpert-Backend/src/database"
	"AgroXpert-Backend/src/middlewares"
	"AgroXpert-Backend/src/routes"
)

func main() {
	configs.InitEnv()
	database.InitMongoConnection()
	app := fiber.New()
	app.Use(middlewares.AccessOriginAnywhere)
	SetupRoutes(app)

	app.Listen(":5000")

	defer database.Db.CloseConnection()
}

func SetupRoutes(app *fiber.App) {
	routes.TestRoutes(app)
	routes.FinalProductionRoutes(app)
	routes.FarmLotRoutes(app)
	routes.EstimatesProductionRoutes(app)
}
