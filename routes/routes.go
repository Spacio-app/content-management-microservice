// routes/routes.go
package routes

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	//"log"
	"github.com/Spacio-app/content-management-microservice/handlers"
	"github.com/Spacio-app/content-management-microservice/models"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/content", validateCourse, handlers.CreateContent)
	// app.Post("/contentPost", validatePost, handlers.CreateContentHandler)
	// app.Post("/contentFile", validateFile, handlers.CreateContentHandler)
	// app.Post("/contentTest", validateTest, handlers.CreateContentHandler)
	app.Get("/content/:id", handlers.GetContentByIDHandler)
	app.Get("/content", handlers.GetAllContentHandler)
	app.Put("/content/:id", handlers.UpdateContentHandler)
	app.Delete("/content/:id", handlers.DeleteContentHandler)
	// Agregar otras rutas para actualizar, eliminar y otras operaciones...
}

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New(validator.WithRequiredStructEnabled())

func validateCourse(c *fiber.Ctx) error {
	var course *models.Courses
	_ = json.Unmarshal([]byte(c.Body()), &course)

	err := Validator.Struct(course)
	if err != nil {
		// log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	return c.Next()
}
