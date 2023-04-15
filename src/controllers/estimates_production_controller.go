package controllers

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllEstimatesProductions(c *fiber.Ctx) error {
	resultEstimatesProduction, err := services.GetAllEstimatesProductions()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if len(resultEstimatesProduction) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Estimates of Production list not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Estimates of Production list successfully",
		Data:    resultEstimatesProduction,
	})
}

func GetOneEstimatesProduction(c *fiber.Ctx) error {
	estimatesProductionID := c.Params("id")
	estimatesProduction, err := services.GetOneEstimatesProduction(estimatesProductionID)

	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Estimates of Production not found, " + err.Error(),
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
		Message: "Estimates of Production successfully",
		Data:    estimatesProduction,
	})
}

func PostNewEstimate(c *fiber.Ctx) error {
	var estimateReq models.ReqEstimate

	if err := c.BodyParser(&estimateReq); err != nil {
		return c.Status(400).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := estimateReq.ValidateEstimate(); err != nil {
		return c.Status(400).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	estimateResult, err := services.CreateEstimate(estimateReq)

	if err != nil {
		return c.Status(400).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(models.Response{
		Success: true,
		Message: "Estimate created successfully",
		Data:    estimateResult,
	})

}
