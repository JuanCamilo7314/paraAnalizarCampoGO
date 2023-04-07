package routes

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/controllers"
)

func FinalProductionRoutes(app *fiber.App) {
	route := app.Group("/final-production")

	route.Get("/", controllers.GetAllFinalProductions)
	route.Get("/:id", controllers.GetOneFinalProduction)

}
