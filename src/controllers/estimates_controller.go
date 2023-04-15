package controllers

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/services"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func PostNewEstimate(c *fiber.Ctx) error {
	var estimateReq models.ReqEstimate

	if err := c.BodyParser(&estimateReq); err != nil {
		return c.Status(400).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := validateEstimate(estimateReq); err != nil {
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

func validateEstimate(estimateReq models.ReqEstimate) error {

	if estimateReq.IdFarmLot == "" {
		return errors.New("IdFarmLot is required")
	}

	if len(estimateReq.TreesAssessed) == 0 || estimateReq.TreesAssessed == nil {
		return errors.New("TreesAssessed is required")
	}

	return nil
}
