package repositories

import (
	"context"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
)

func CreateCourse(content interface{}) error {
	collection := utils.GetCollection("Content")
	log.Printf("Content: %v\n", content)
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el curso: %v\n", err)
	}
	return err
}

// get all courses from content collection
func GetAllCourses() ([]models.Courses, error) {
	collection := utils.GetCollection("Content")
	var courses []models.Courses
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Printf("Error al obtener los cursos: %v\n", err)
		return courses, err
	}
	err = cursor.All(context.Background(), &courses)
	if err != nil {
		log.Printf("Error al obtener los cursos: %v\n", err)
		return courses, err
	}
	return courses, nil
}

// // update course
// func UpdateCourse(id string, content interface{}) error {
// 	collection := utils.GetCollection("Content")
// 	_, err := collection.UpdateOne(context.Background(), id, content)
// 	if err != nil {
// 		log.Printf("Error al actualizar el curso: %v\n", err)
// 	}
// 	return err
// }
