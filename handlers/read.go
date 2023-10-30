// handlers/read_handler.go
package handlers

import (
	"strconv"

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el contenido",
		})
	}

	return c.JSON(content)
}
func GetContentFeedHandler(c *fiber.Ctx) error {
	//obtener parametros de paginacion y ordenamiento
	page := c.Query("page", "1") //pagina por defecto 1
	limit := c.Query("limit", "10")

	//calcula el numero de documentos a saltar
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
	skip := (pageInt - 1) * limitInt
	//realizar la consulta
	content, err := services.GetContentFeedOrderByDate(skip, limitInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el feed de contenido",
		})
	}

	return c.JSON(content)
}
