package handlers

import (
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
)

// Initialize the Cloudinary client using environment variables

// validar datos de entrada
func CreateCourse(c *fiber.Ctx) error {
	// Crear estructura de contenido de datos
	content := domain.CourseReq{}

	// Analizar el cuerpo de la solicitud y almacenar los datos en la estructura
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}

	// Verificar si se proporcionaron videos
	if len(content.Videos) > 0 {
		isVideo := true
		// Procesar y cargar archivos
		secureURLs, publicIDs, miniatureURL, err := utils.ProcessUploadedFiles(c, "VideosURL", isVideo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error al procesar archivos",
			})
		}

		// Asignar los resultados a los elementos individuales del slice
		for i, secureURL := range secureURLs {
			content.Videos.URL[i] = secureURL
		}
		for i, publicID := range publicIDs {
			content.PublicIDCloudinary[i] = publicID
		}
		// miniatureURL es una string simple
		content.Miniature = miniatureURL
	}

	log.Println("Creando un nuevo curso...")

	if err := services.CreateCourse(content); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al crear el curso",
		})
	}
	return c.JSON(content)
}

func GetAllCoursesHandler(c *fiber.Ctx) error {
	content, err := services.GetAllCourses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener los cursos",
		})
	}
	return c.JSON(content)
}

func UpdateCourseHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	content := domain.CourseReq{}
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}
	err := services.UpdateCourse(id, content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al actualizar el curso",
		})
	}
	return c.JSON(content)
}
