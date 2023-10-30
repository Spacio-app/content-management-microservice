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

func GetContentByID(id primitive.ObjectID) (models.GenericContent, error) {
	return repositories.GetContentByID(id)
}
func GetAllContent() ([]models.GenericContent, error) {
	content, err := repositories.GetAllContent()
	if err != nil {
		return nil, errors.New("error al obtener el contenido")
	}
	return content, nil
}
func GetContentFeedOrderByDate(skip, limit int) ([]models.GenericContent, error) {
	//
	return repositories.GetContentFeedOrderByDate(skip, limit)
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

// updates
func UpdateCourse(id string, content domain.CourseReq) error {
	// Convertir la cadena de texto a ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	// Verificar si existe el contenido
	_, err = repositories.GetContentByID(objectID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.UpdateCourse(objectID, content)
}

func UpdatePost(id string, content domain.PostReq) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}
	_, err = repositories.GetContentByID(objectID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.UpdatePost(objectID, content)
}
func UpdateFile(id string, content domain.FileReq) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}
	_, err = repositories.GetContentByID(objectID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.UpdateFile(objectID, content)
}
func UpdateTest(id string, content domain.TestReq) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("ID inválido")
	}
	_, err = repositories.GetContentByID(objectID)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.UpdateTest(objectID, content)
}

// DeleteContentByID elimina un registro de contenido por ID del repositorio
func DeleteContentByID(id primitive.ObjectID) error {
	//verificar si existe el contenido
	_, err := repositories.GetContentByID(id)
	if err != nil {
		return errors.New("el contenido no existe")
	}
	return repositories.DeleteContentByID(id)
}

// Continuar con las funciones para actualizar y eliminar contenido...
