package routes

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/controllers"
)

func EstimatesProductionRoutes(app *fiber.App) {
	route := app.Group("/estimates-production")

	route.Get("/", controllers.GetAllEstimatesProductions)
	route.Get("/:id", controllers.GetOneEstimatesProduction)
}
