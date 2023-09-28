// services/content_service.go
package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// llamar repos, clientes, apis externas, para guardar, actualizar, eliminar.
func CreateCourse(content domain.CourseReq) error {
	fmt.Println("contentasds", content)
	log.Println("Guardando curso en la base de datos...")

	return repositories.CreateCourse(content)
}
func CreatePost(content domain.PostReq) error {
	return repositories.CreatePost(content)
}
func CreateFile(content domain.FileReq) error {
	return repositories.CreateFile(content)
}
func CreateTest(content domain.TestReq) error {
	return repositories.CreateTest(content)
}

func GetContentByID(id primitive.ObjectID) (*models.Courses, error) {
	return repositories.GetContentByID(id)
}
func GetAllContent() ([]models.GenericContent, error) {
	return repositories.GetAllContent()
}
func GetAllCourses() ([]models.Courses, error) {
	courses, err := repositories.GetAllCourses()
	if err != nil {
		return nil, errors.New("error al obtener los cursos")
	}
	return courses, nil
}
func GetAllPosts() ([]models.Posts, error) {
	return repositories.GetAllPosts()
}
func GetAllFiles() ([]models.Files, error) {
	return repositories.GetAllFiles()
}
func GetAllTests() ([]models.Tests, error) {
	return repositories.GetAllTests()
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
