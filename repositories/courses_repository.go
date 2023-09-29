package repositories

import (
	"context"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCourse(content domain.CourseReq) error {
	collection := utils.GetCollection("Content")
	content.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar
	log.Printf("Content: %v\n", content)
	log.Println("Insertando datos en la base de datos...")
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el curso: %v\n", err)
	}
	return err
}

// get all courses from content collection
func GetAllCourses() ([]models.Courses, error) {
	collection := utils.GetCollection("Content")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var course []models.Courses
	if err := cursor.All(context.Background(), &course); err != nil {
		return nil, err
	}
	return course, nil
}

// update course
func UpdateCourse(id primitive.ObjectID, content domain.CourseReq) error {
	collection := utils.GetCollection("Content")
	content.BeforeUpdate() // Actualiza updatedAt antes de actualizar
	//mantener createdAt original
	filter := bson.M{"_id": id}
	update := bson.M{"$set": content}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
