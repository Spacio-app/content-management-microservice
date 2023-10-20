// routes/routes.go
package routes

import (
	//"encoding/json"

	//"github.com/go-playground/validator/v10"
	"fmt"
	"net/http"

	"github.com/Spacio-app/content-management-microservice/handlers"
	"github.com/gofiber/fiber/v2"
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

	app.Patch("/contentCourse/:id", handlers.UpdateCourseHandler)
	app.Put("/contentPost/:id", handlers.UpdatePostHandler)
	app.Put("/contentFile/:id", handlers.UpdateFileHandler)
	app.Put("/contentTest/:id", handlers.UpdateTestHandler)

	app.Get("/Content/:id", handlers.GetContentByIDHandler)
	app.Get("/Content", handlers.GetAllContentHandler)

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

//		err := Validator.Struct(course)
//		if err != nil {
//			// log.Error(err)
//			return c.Status(fiber.StatusBadRequest).JSON(err)
//		}
//		return c.Next()
//	}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Tamaño máximo del archivo (10 MB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file") // Nombre del campo del formulario
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Aquí puedes guardar el archivo en Cloudinary o realizar cualquier otra operación
	// con el archivo cargado.

	fmt.Fprintf(w, "Archivo cargado correctamente")
}
