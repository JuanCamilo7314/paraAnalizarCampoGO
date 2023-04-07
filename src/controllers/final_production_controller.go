package controllers

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/services"

	"github.com/gofiber/fiber/v2"
)

func GetAllFinalProductions(c *fiber.Ctx) error {
	resultFinalProduction, err := services.GetAllFinalProductions()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: "Error getting Final Production list",
		})
	}

	if len(resultFinalProduction) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Final Production list not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Final Production list successfully",
		Data:    resultFinalProduction,
	})
}

/*func GetOneFinalProduction(c *fiber.Ctx) error {
	idFinalProduc := c.Params("id")
	fmt.Println("ID: ", id)
	return c.SendString("Final Production: " + id)
}*/
