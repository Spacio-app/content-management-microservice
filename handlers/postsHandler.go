package handlers

import (
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {

	content := domain.PostReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}

	//Procesar y cargar archivos
	if secureURL, publicID, err := utils.ProcessUploadedFiles(c, "ImagesURL"); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al procesar archivos",
		})
	} else {
		content.ImagesURL = secureURL
		content.PublicIDCloudinary = publicID
		content.Miniature = secureURL[0]
	}

	log.Println("Creando un nuevo post...")

	//enviar a servicio
	err := services.CreatePost(content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al crear el post",
		})
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

// update post
func UpdatePostHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	content := domain.PostReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	err := services.UpdatePost(id, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el post",
		})
	}
	return c.JSON(content)
}
