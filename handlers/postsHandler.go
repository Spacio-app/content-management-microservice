package handlers

import (
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {

	content := domain.PostReq{}

	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}

	return c.JSON(content)

}

// get all posts
func GetAllPostsHandler(c *fiber.Ctx) error {
	content, err := services.GetAllPosts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener los posts",
		})
	}

	return c.JSON(content)
}
