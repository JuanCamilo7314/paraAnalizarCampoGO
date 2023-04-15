package routes

import (
	"AgroXpert-Backend/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func EstimatesProductionRoutes(app *fiber.App) {
	route := app.Group("/estimates-production")

	route.Get("/", controllers.GetAllEstimatesProductions)
	route.Get("/:id", controllers.GetOneEstimatesProduction)
	route.Post("/harvest", controllers.EstimatesPerHarvest)
	route.Post("/", controllers.PostNewEstimate)
}
