package routes

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/controllers"
)

func HarvestRoutes(app *fiber.App) {
	route := app.Group("/harvest")

	route.Get("/", controllers.GetAllHarvests)
	route.Get("/:id", controllers.GetOneHarvest)
	route.Get("/historic/:idFarmLot", controllers.GetHistoricHarvestEsimation)
	route.Post("/", controllers.CreateHarvest)
}
