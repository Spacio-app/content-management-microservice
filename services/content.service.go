// services/content_service.go
package services

import (
	"github.com/Spacio-app/content-management-microservice/models"
	"github.com/Spacio-app/content-management-microservice/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateContent(content interface{}) error {

	return repositories.CreateCourse(content)

}

func GetContentByID(id primitive.ObjectID) (*models.Courses, error) {
	return repositories.GetContentByID(id)
}
func GetAllContent() ([]models.Courses, error) {
	return repositories.GetAllContent()
}

// UpdateContentByID actualiza un registro de contenido por ID en el repositorio
func UpdateContentByID(id primitive.ObjectID, updatedContent models.Courses) error {
	return repositories.UpdateContentByID(id, updatedContent)
}

// DeleteContentByID elimina un registro de contenido por ID del repositorio
func DeleteContentByID(id primitive.ObjectID) error {
	return repositories.DeleteContentByID(id)
}

// Continuar con las funciones para actualizar y eliminar contenido...
