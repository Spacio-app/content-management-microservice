package handlers

import (
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTest(c *fiber.Ctx) error {
	content := domain.TestReq{}
	content.ID = primitive.NewObjectID().Hex() // Generar un ID único
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	//enviar a servicio
	err := services.CreateTest(content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al crear el test",
		})
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
