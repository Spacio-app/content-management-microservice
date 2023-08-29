// handlers/read_handler.go
package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/Spacio-app/content-management-microservice/services"
)

func GetContentByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	content, err := services.GetContentByID(id)
	if err != nil {
		log.Println("Error al obtener el contenido por ID en el handler:", err)
		return err
	}

	return c.JSON(content)
}
