package repositories

import (
	"context"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
)

func CreateFile(content domain.FileReq) error {
	collection := utils.GetCollection("Content")
	content.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el curso: %v\n", err)
	}
	return err
}

func GetAllFiles() ([]models.Files, error) {
	var files []models.Files
	collection := utils.GetCollection("Content")
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Printf("Error al obtener los cursos: %v\n", err)
		return nil, err
	}
	err = cursor.All(context.Background(), &files)
	if err != nil {
		log.Printf("Error al obtener los cursos: %v\n", err)
		return nil, err
	}
	return files, nil
}

//  update file
// func UpdateFile(id string, content interface{}) error {
// 	collection := utils.GetCollection("Content")
// 	_, err := collection.UpdateOne(context.Background(), id, content)
// 	if err != nil {
// 		log.Printf("Error al actualizar el curso: %v\n", err)
// 	}
// 	return err
// }
