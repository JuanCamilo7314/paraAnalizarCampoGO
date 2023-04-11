package controllers

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllFinalProductions(c *fiber.Ctx) error {
	resultFinalProduction, err := services.GetAllFinalProductions()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
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

func GetOneFinalProduction(c *fiber.Ctx) error {
	finalProductionID := c.Params("id")
	finalProduction, err := services.GetOneFinalProduction(finalProductionID)

	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Final Production not found, " + err.Error(),
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Final Production successfully",
		Data:    finalProduction,
	})
}
