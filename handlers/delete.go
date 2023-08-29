// handlers/delete_handler.go
package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/Spacio-app/content-management-microservice/services"
)

func DeleteContentByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteContentByID(id)
	if err != nil {
		log.Println("Error al eliminar el contenido por ID en el handler:", err)
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
