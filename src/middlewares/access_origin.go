package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AccessOriginAnywhere(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	if c.Method() == "OPTIONS" {
		return c.SendStatus(200)
	}
	return c.Next()
}
