package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAllFinalProductions(c *fiber.Ctx) error {
	//return c.JSON(services.GetAllFinalProductions())
	return c.SendString("All Final Productions")
}

func GetFinalProduction(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("ID: ", id)
	return c.SendString("Final Production: " + id)
}
