// handlers/read_handler.go
package handlers

import (
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetContentByIDHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Convertir el ID en formato string a un tipo ObjectId
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	content, err := services.GetContentByID(objectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el contenido por ID",
		})
	}

	return c.JSON(content)
}
func GetAllContentHandler(c *fiber.Ctx) error {
	content, err := services.GetAllContent()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el contenido",
		})
	}

	return c.JSON(content)
}
