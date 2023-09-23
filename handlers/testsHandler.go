package handlers

import (
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
)

func CreateTest(c *fiber.Ctx) error {
	content := domain.TestReq{}

	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}

	return c.JSON(content)
}

// get all tests
func GetAllTestsHandler(c *fiber.Ctx) error {
	content, err := services.GetAllTests()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener los tests",
		})
	}

	return c.JSON(content)
}
