package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/samuel", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World from Samuel!")
	})

	app.Listen(":5000")
}
