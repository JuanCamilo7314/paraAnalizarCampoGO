package controllers

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/services"

	"github.com/gofiber/fiber/v2"
)

func GetGraphsHarvest(c *fiber.Ctx) error {
	FarmLotID := c.Params("idFarmLot")

	dataGraphHarvest, err := services.GetGraphsHarvest(FarmLotID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Graphs Harvest",
		Data:    dataGraphHarvest,
	})
}
