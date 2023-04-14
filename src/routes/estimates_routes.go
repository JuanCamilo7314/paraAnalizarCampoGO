package routes

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/controllers"
)

func EstimatesRoutes(app *fiber.App) {
	route := app.Group("/estimate")

	route.Post("/", controllers.PostNewEstimate)
}
