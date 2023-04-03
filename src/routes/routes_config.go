package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/samuel", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Samuel!")
	})

	app.Get("/cristian", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Cristian this is a Software project!")
	})

	app.Get("/yorman", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Yorman!")
	})

	app.Get("/camila", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Camila!")
	})
}
