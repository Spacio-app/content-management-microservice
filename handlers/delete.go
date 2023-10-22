package handlers

import (
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteContentHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	// Obtén el contenido que se va a eliminar
	content, err := services.GetContentByID(objectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el contenido por ID",
		})
	}
	//verificar si el contenido es un curso
	isCourse := content.Videos != nil
	//verificar si el contenido es un archivo
	isFile := content.FilesURL != nil
	//verificar si el contenido es un Post
	isPost := content.ImagesURL != nil
	if isCourse {
		for _, video := range content.Videos {
			err = utils.DeleteContentFromCloudinary(video.PublicIDCloudinary, true)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Error al eliminar el archivo de Cloudinary",
				})
			}
		}
	}
	if isFile {
		for _, file := range content.FilesURL {
			err = utils.DeleteContentFromCloudinary(file.PublicIDCloudinary, false)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Error al eliminar el archivo de Cloudinary",
				})
			}
		}
	}
	if isPost {
		for _, image := range content.ImagesURL {
			err = utils.DeleteContentFromCloudinary(image.PublicIDCloudinary, false)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Error al eliminar el archivo de Cloudinary",
				})
			}
		}
	}

	// Elimina el contenido de la base de datos
	err = services.DeleteContentByID(objectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al eliminar el contenido por ID",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Registro eliminado con éxito",
	})
}
