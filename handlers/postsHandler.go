package handlers

import (
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	content := domain.PostReq{}

	// Acceder a los campos del formulario
	title := c.FormValue("title")
	description := c.FormValue("description")
	// Asignar los valores de título y descripción
	content.Title = title
	content.Description = description
	user, error := getUserHeader(c)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el usuario",
		})
	}
	author := domain.AuthorReq{
		Name:  user.Name,
		Photo: user.Image,
	}

	content.Author = author
	// Procesar las imágenes
	isVideo := false

	images := []domain.ImageURLReq{}
	for i := 0; ; i++ {
		imageKey := fmt.Sprintf("imagesURL[%d][imageURL]", i)
		image, err := c.FormFile(imageKey)
		if err != nil {
			// No more images to process
			if image == nil {
				break
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al procesar archivos",
			})
		}

		secureURL, publicID, _, _, err := utils.ProcessUploadedFiles(c, image, isVideo, nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al procesar archivos",
			})
		}

		imageInfo := domain.ImageURLReq{
			ImageURL:           secureURL,
			PublicIDCloudinary: publicID,
		}

		images = append(images, imageInfo)
	}

	// Asignar los resultados de procesar las imágenes
	content.ImagesURL = images

	log.Println("Creando un nuevo post...")

	// Enviar a servicio
	if err := services.CreatePost(content); err != nil {
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
