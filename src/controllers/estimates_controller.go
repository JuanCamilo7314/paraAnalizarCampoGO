package controllers

import (
	"AgroXpert-Backend/src/models"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func PostNewEstimate(c *fiber.Ctx) error {
	var estimate models.EstimateModel
	fmt.Println(reflect.TypeOf(c.Body()))

	if err := c.BodyParser(&estimate); err != nil {
		return c.Status(400).JSON(models.Response{
			Success: false,
			Message: err.Error(),
			Data:    c.Body(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Estimate created successfully",
		"data":    estimate,
	})
}
