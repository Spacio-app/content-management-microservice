package repositories

import (
	"context"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
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
	cursor, err := collection.Find(context.Background(), nil)
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
