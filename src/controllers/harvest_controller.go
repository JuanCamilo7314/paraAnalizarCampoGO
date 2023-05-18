package controllers

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllHarvests(c *fiber.Ctx) error {
	resultHarvest, err := services.GetAllHarvests()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if len(resultHarvest) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Harvest list not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Harvest list successfully",
		Data:    resultHarvest,
	})
}

func GetOneHarvest(c *fiber.Ctx) error {
	HarvestID := c.Params("id")
	Harvest, err := services.GetOneHarvest(HarvestID)

	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Harvest not found, " + err.Error(),
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
		Message: "Harvest successfully",
		Data:    Harvest,
	})
}

func CreateHarvest(c *fiber.Ctx) error {
	var harvestReq models.CreateHarvest

	if err := c.BodyParser(&harvestReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := harvestReq.ValidateHarvest(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	harvestResult, err := services.CreateHarvest(harvestReq)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Harvest successfully",
		Data:    harvestResult,
	})
}

func GetHistoricHarvestEsimation(c *fiber.Ctx) error {
	FarmLotID := c.Params("idFarmLot")
	historicHarvest, err := services.GetHistoricHarvestEsimation(FarmLotID)
	//fmt.Printf("%+v", historicHarvest)

	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Historic Harvest not found, " + err.Error(),
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
		Message: "Historic Harvest successfully",
		Data:    historicHarvest,
	})
}
