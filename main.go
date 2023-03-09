package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/samuel", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Samuel!")
	})

	app.Get("/cristian", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Cristian this is a Software project!")
	})

	app.Listen(":5000")
}
