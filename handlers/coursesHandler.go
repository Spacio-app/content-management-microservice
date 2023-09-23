package handlers

import (
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
)

func CreateCourse(c *fiber.Ctx) error {

	content := domain.CourseReq{}
	fmt.Println("content", content)
	if err := c.BodyParser(&content); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return err
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
