package handlers

import (
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateContentHandler maneja la solicitud para actualizar un registro de contenido por ID
func UpdateContentHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	var updatedContent models.Courses
	if err := c.BodyParser(&updatedContent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	err = services.UpdateContentByID(objectID, updatedContent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el contenido por ID",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Registro actualizado con éxito",
	})
}
