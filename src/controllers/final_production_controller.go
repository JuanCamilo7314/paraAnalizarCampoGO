package controllers

import (
	"fmt"

	"AgroXpert-Backend/src/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllFinalProductions(c *fiber.Ctx) error {
	resultFinalProduction, err := services.GetAllFinalProductions()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
			"succes":  false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Final Production list successfully",
		"succes":  true,
		"data":    resultFinalProduction,
	})
}

func GetFinalProduction(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("ID: ", id)
	return c.SendString("Final Production: " + id)
}
