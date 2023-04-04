package routes

import (
	"AgroXpert-Backend/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupTestRoutes(app *fiber.App) {
	route := app.Group("/test")

	route.Get("/samuel", controllers.GetSamuel)
	route.Get("/cristian", controllers.GetCristian)
	route.Get("/yorman", controllers.GetYorman)
	route.Get("/camila", controllers.GetCamila)
	route.Get("/getDataTest", controllers.GetDataTest)
}
