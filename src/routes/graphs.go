package routes

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/controllers"
)

func GraphsRoutes(app *fiber.App) {
	route := app.Group("/graphs")
	route.Get("/:idFarmLot", controllers.GetGraphsHarvest)
}
