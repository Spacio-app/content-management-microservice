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
	content.Title = title
	content.Description = description

	announcement := c.FormValue("createAnnouncement")
	fmt.Println("announcement", announcement)
	content.CreateAnnouncement = announcement == "true"

	// user, error := getUserHeader(c)
	// if error != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "Error al obtener el usuario",
	// 	})
	// }
	author := domain.AuthorReq{
		Name:  "user.Name",
		Photo: "user.Image",
	}

	URLMiniature, publicIDMiniature, err := uploadMiniature(c)
	if err != nil {
		content.Miniature = ""
	}

	content.Miniature = URLMiniature
	content.PublicIDMiniature = publicIDMiniature

	content.Author = author
	content.FilesURL = []domain.FileURLReq{}
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
		secureURL, publicID, err := utils.UploadRaws(c, file)
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
	if content.CreateAnnouncement {
		announcement := createAnnouncementFromFile(content)
		if err := services.CreateAnnouncement(announcement); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al crear el anuncio",
			})
		}
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
func createAnnouncementFromFile(content domain.FileReq) domain.FeedReq {
	announcement := domain.FeedReq{
		Title:       content.Title,
		Description: "Se ha creado un nuevo documento: " + content.Title + "\n" + content.Description,
		Author:      content.Author,
	}
	return announcement
}
