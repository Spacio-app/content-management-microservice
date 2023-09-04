package repositories

import (
	"context"
	"log"

	"github.com/Spacio-app/content-management-microservice/models"
	"github.com/Spacio-app/content-management-microservice/utils"
)

func CreateCourse(content interface{}) error {
	collection := utils.GetCollection("content")
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el curso: %v\n", err)
	}
	return err
}
func CreatePost(content *models.Posts) error {
	collection := utils.GetCollection("content")
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el post: %v\n", err)
	}
	return err
}
