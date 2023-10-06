package handlers

import (
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Obtén el contenido que se va a eliminar
	content, err := services.GetContentByID(objectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el contenido por ID",
		})
	}

	// Elimina el contenido de la base de datos
	err = services.DeleteContentByID(objectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar el contenido por ID",
		})
	}

	// Llama a la función para eliminar el archivo de Cloudinary
	err = utils.DeleteContentFromCloudinary(content.PublicIDCloudinary)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar el archivo de Cloudinary",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Registro eliminado con éxito",
	})
}
