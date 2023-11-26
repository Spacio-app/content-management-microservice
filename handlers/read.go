// handlers/read_handler.go
package handlers

import (
	"fmt"
	"strconv"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetContentByIDHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// Convertir el ID en formato string a un tipo ObjectId
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	content, err := services.GetContentByID(objectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el contenido por ID",
		})
	}

	return c.JSON(content)
}
func GetAllContentHandler(c *fiber.Ctx) error {
	content, err := services.GetAllContent()
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el contenido",
		})
	}

	return c.JSON(content)
}
func GetContentFeedHandler(c *fiber.Ctx) error {
	// Obtener los parámetros de paginación y ordenamiento
	page := c.Query("page", "1") // Página predeterminada: 1
	limit := c.Query("limit", "5")

	// Calcular el número de documentos a saltar
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error al obtener el contenido",
		})
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error al obtener el contenido",
		})
	}

	// Calcular el valor de "skip" basado en la página y el límite
	skip := (pageInt - 1) * limitInt

	// Realizar la consulta usando el mismo servicio
	content, err := services.GetContentFeedOrderByDate(skip, limitInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el feed de contenido",
		})
	}

	// Verificar si no hay más contenido
	if len(content) == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": "No hay más contenido de feed disponible.",
		})
	}

	// Devolver los datos como respuesta JSON
	return c.JSON(content)
}

func GetContentByAuthorHandler(c *fiber.Ctx) error {
	user, error := getUserHeader(c)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el usuario",
		})
	}
	authorHeader := domain.AuthorReq{
		Name:  user.Name,
		Photo: user.Image,
		Email: user.Email,
	}
	//correo del autor
	author := authorHeader.Email

	// Llama a tu función en el servicio que obtiene el contenido por autor
	content, err := services.GetContentByAuthor(author)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el contenido del autor",
		})
	}

	return c.JSON(content)
}
