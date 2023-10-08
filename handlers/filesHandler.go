package handlers

import (
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateFile(c *fiber.Ctx) error {
	content := domain.FileReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	isVideo := false
	//Procesar y cargar archivos
	if secureURL, publicID, miniature, err := utils.ProcessUploadedFiles(c, "FilesURL", isVideo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al procesar archivos",
		})
	} else {
		content.FilesURL = secureURL
		content.PublicIDCloudinary = publicID
		content.Miniature = miniature
	}

	log.Println("Creando un nuevo file...")

	//enviar a servicio
	err := services.CreateFile(content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al crear el archivo",
		})
	}

	return c.JSON(content)
}

func GetAllFilesHandler(c *fiber.Ctx) error {
	content, err := services.GetAllFiles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener los archivos",
		})
	}

	return c.JSON(content)
}

func UpdateFileHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	content := domain.FileReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	err := services.UpdateFile(id, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el archivo",
		})
	}
	return c.JSON(content)
}
