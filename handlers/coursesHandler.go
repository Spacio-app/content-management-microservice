package handlers

import (
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
)

// Initialize the Cloudinary client using environment variables

// validar datos de entrada
func CreateCourse(c *fiber.Ctx) error {
	// crear estructra de contenido de datos
	content := domain.CourseReq{}
	fmt.Println("content", c)
	log.Println("Creando un nuevo curso...")
	// Analizar el cuerpo de la solicitud y almacenar los datos en la estructura
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
	}

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
