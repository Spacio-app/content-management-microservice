// routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	//"log"

	"github.com/Spacio-app/content-management-microservice/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/content", handlers.CreateContentHandler)
	app.Get("/content/:id", handlers.GetContentByIDHandler)
	app.Put("/content/:id", handlers.UpdateContentHandler)
	app.Delete("/content/:id", handlers.DeleteContentByIDHandler)
	// Agregar otras rutas para actualizar, eliminar y otras operaciones...
}
