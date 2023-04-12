package routes

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/controllers"
)

func FarmLotRoutes(app *fiber.App) {
	route := app.Group("/farm-lot")

	route.Get("/", controllers.GetAllFarmLots)
	route.Get("/:id", controllers.GetOneFarmLot)
}
