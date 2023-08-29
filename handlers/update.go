// handlers/update_handler.go
package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/Spacio-app/content-management-microservice/models"
	"github.com/Spacio-app/content-management-microservice/services"
)

func UpdateContentHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	var updatedContent models.Content
	if err := c.BodyParser(&updatedContent); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}

	updatedContent.ID = id

	if err := services.UpdateContent(&updatedContent); err != nil {
		log.Println("Error al actualizar el contenido en el handler:", err)
		return err
	}

	return c.JSON(updatedContent)
}
