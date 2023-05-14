package controllers

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllFarmLots(c *fiber.Ctx) error {
	resultFarmLot, err := services.GetAllFarmLots()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if len(resultFarmLot) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Farm Lot list not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Farm Lot list successfully",
		Data:    resultFarmLot,
	})
}

func GetOneFarmLot(c *fiber.Ctx) error {
	finalProductionID := c.Params("id")
	finalProduction, err := services.GetOneFarmLot(finalProductionID)

	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Farm Lot not found, " + err.Error(),
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
		Message: "Farm Lot successfully",
		Data:    finalProduction,
	})
}

func CreateFarmLot(c *fiber.Ctx) error {
	var farmLotReq models.FarmLotReq

	if err := c.BodyParser(&farmLotReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	// if err := farmLotReq.ValidateFarmLot(); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(models.Response{
	// 		Success: false,
	// 		Message: err.Error(),
	// 	})
	// }

	farmLotResult, err := services.CreateFarmLot(farmLotReq)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Farm Lot successfully",
		Data:    farmLotResult,
	})
}
