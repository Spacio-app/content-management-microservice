// handlers/create_handler.go
package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/Spacio-app/content-management-microservice/services"
)

func CreateContent(c *fiber.Ctx) error {
	var content interface{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	if err := services.CreateContent(&content); err != nil {
		log.Println("Error al crear el contenido en el handler:", err)
		return err
	}
	return c.JSON(content)
}
