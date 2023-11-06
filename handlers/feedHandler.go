package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
)

func CreateFeed(c *fiber.Ctx) error {
	content := domain.FeedReq{}
	//bodyparser
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	UserHeader := c.Get("User")

	var user User

	if err := json.Unmarshal([]byte(UserHeader), &user); err != nil {
		fmt.Println("Error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al procesar el usuario",
		})
	}

	content.Author = domain.AuthorReq{
		Name:  user.Name,
		Photo: user.Image,
	}
	// AuthorID := c.Locals("userID").(string)     // Suponiendo que userID es un string
	// AuthorPhoto := c.Locals("userPhoto").(string)

	//generar comentarios vacios
	content.Comments = []domain.FeedCommentsReq{}

	//llamar al servicio
	if err := services.CreateFeed(content); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al crear el curso",
		})
	}
	return c.JSON(content)
}
func UpdatePostComments(c *fiber.Ctx) error {
	// Obtener postID del parámetro o cuerpo de la solicitud
	postID := c.Params("postID")

	fmt.Println("postID", postID)
	// Obtener los nuevos datos del array de comentarios del cuerpo de la solicitud
	// var updatedComments domain.FeedCommentsReq

	var comment1 domain.FeedCommentsReq
	if err := c.BodyParser(&comment1); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	UserHeader := c.Get("User")
	var user User

	if err := json.Unmarshal([]byte(UserHeader), &user); err != nil {
		fmt.Println("Error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al procesar el usuario",
		})

	}

	// Crear un nuevo comentario con el autor adecuado
	authorReq := domain.AuthorReq{
		Name:  user.Name,
		Photo: user.Image,
	}

	comment := domain.FeedCommentsReq{
		Author:  authorReq,
		Comment: comment1.Comment,
		// Asignar otros campos de comentario si es necesario
	}

	// Agregar el nuevo comentario a la lista de comentarios actualizados

	// Llamar a un servicio para actualizar los comentarios del post
	if err := services.UpdatePostComments(postID, comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar los comentarios del post",
		})
	}
	return c.JSON(comment)
}

func GetAllFeedsHandler(c *fiber.Ctx) error {
	content, err := services.GetAllFeeds()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener los archivos",
		})
	}
	return c.JSON(content)
}

//funcion para generar feed a partir de los contenidos si el usuario lo desea

// func generateFeedFromContent(content domain.GenericContent) {

// }
