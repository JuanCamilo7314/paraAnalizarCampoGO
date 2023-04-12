package controllers

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/services"
)

func GetSamuel(c *fiber.Ctx) error {
	return c.SendString("Hello, World from Samuel!")
}

func GetCristian(c *fiber.Ctx) error {
	return c.SendString("Hello, World from Cristian this is a Software project!")
}

func GetYorman(c *fiber.Ctx) error {
	return c.SendString("Hello, World from Yorman XD XD!")
}

func GetCamila(c *fiber.Ctx) error {
	return c.SendString("Hello, World from Camila!")
}

func GetDataTest(c *fiber.Ctx) error {
	services.GetDataTest()
	return c.SendString("Mira la consola, pueden salir cositas")
}
