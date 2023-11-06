package handlers

import (
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateFile(c *fiber.Ctx) error {
	content := domain.FileReq{}

	title := c.FormValue("title")
	description := c.FormValue("description")
	author := c.FormValue("author")
	content.Title = title
	content.Description = description
	content.Author = author
	content.FilesURL = []domain.FileURLReq{}
	isVideo := false
	for i := 0; ; i++ {
		fileKey := fmt.Sprintf("filesURL[%d][fileURL]", i)
		file, err := c.FormFile(fileKey)
		if err != nil {
			// No more images to process
			if file == nil {
				break
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al procesar archivos",
			})
		}

		// Procesar y cargar archivos
		secureURL, publicID, _, _, err := utils.ProcessUploadedFiles(c, file, isVideo, nil)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al procesar archivos",
			})
		}

		// Crear un nuevo elemento en la lista
		newFileURL := domain.FileURLReq{
			FileURL:            secureURL,
			PublicIDCloudinary: publicID,
		}

		// Agregar el nuevo elemento a la lista
		content.FilesURL = append(content.FilesURL, newFileURL)

		// Actualizar el campo Miniature (deberías considerar cómo gestionar este campo)

		log.Printf("Creando un nuevo archivo (índice %d)...\n", i)
	}

	// Enviar a servicio
	if err := services.CreateFile(content); err != nil {
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
