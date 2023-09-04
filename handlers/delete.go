package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Spacio-app/content-management-microservice/services"
)

// DeleteContentHandler maneja la solicitud para eliminar un registro de contenido por ID
func DeleteContentHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	err = services.DeleteContentByID(objectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar el contenido por ID",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Registro eliminado con éxito",
	})
}
