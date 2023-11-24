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

func CreateTest(content domain.TestReq) error {
	collection := utils.GetCollection("Content")
	content.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el curso: %v\n", err)
	}
	return err
}

func GetAllTests() ([]models.Tests, error) {
	var tests []models.Tests
	collection := utils.GetCollection("Content")
	filter := bson.M{"contenttype": "test"}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error al obtener los cursos: %v\n", err)
		return nil, err
	}
	err = cursor.All(context.Background(), &tests)
	if err != nil {
		log.Printf("Error al obtener los cursos: %v\n", err)
		return nil, err
	}
	return tests, nil
}

// update test
func UpdateTest(id primitive.ObjectID, content domain.TestReq) error {
	collection := utils.GetCollection("Content")
	content.BeforeUpdate() // Actualiza updatedAt antes de actualizar
	filter := bson.M{"_id": id}
	update := bson.M{"$set": content}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
func SaveTestResult(content domain.TestResultReq) error {
	collection := utils.GetCollection("TestResult")
	content.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el curso: %v\n", err)
	}
	return err
}
func GetTestResult(contentID primitive.ObjectID, email string) (float64, error) {
	collection := utils.GetCollection("TestResult")
	filter := bson.M{"content_id": contentID, "author.email": email}
	var result models.TestResult
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Calification, nil
}
