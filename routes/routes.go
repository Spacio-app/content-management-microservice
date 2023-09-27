// routes/routes.go
package routes

import (
	//"encoding/json"

	//"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/Spacio-app/content-management-microservice/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/contentCourse", handlers.CreateCourse)
	app.Post("/contentPost", handlers.CreatePost)
	app.Post("/contentFile", handlers.CreateFile)
	app.Post("/contentTest", handlers.CreateTest)
	app.Get("/contentCourse/", handlers.GetAllCoursesHandler)
	app.Get("/contentPost/", handlers.GetAllPostsHandler)
	app.Get("/contentFile/", handlers.GetAllFilesHandler)
	app.Get("/contentTest/", handlers.GetAllTestsHandler)

	app.Get("/Content/:id", handlers.GetContentByIDHandler)
	app.Get("/Content", handlers.GetAllContentHandler)
	app.Put("/Content/:id", handlers.UpdateContentHandler)
	app.Delete("/Content/:id", handlers.DeleteContentHandler)
	// Agregar otras rutas para actualizar, eliminar y otras operaciones...
}

type IError struct {
	Field string
	Tag   string
	Value string
}

// var Validator = validator.New(validator.WithRequiredStructEnabled())

// func validateCourse(c *fiber.Ctx) error {
// 	var course *models.Courses
// 	fmt.Println("course", course)
// 	_ = json.Unmarshal([]byte(c.Body()), &course)

// 	err := Validator.Struct(course)
// 	if err != nil {
// 		// log.Error(err)
// 		return c.Status(fiber.StatusBadRequest).JSON(err)
// 	}
// 	return c.Next()
// }
