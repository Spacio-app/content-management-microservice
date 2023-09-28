package repositories

import (
	"context"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
)

func CreatePost(content domain.PostReq) error {
	collection := utils.GetCollection("Content")
	content.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el post: %v\n", err)
	}
	return err
}

func GetAllPosts() ([]models.Posts, error) {
	var posts []models.Posts
	collection := utils.GetCollection("Content")
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Printf("Error al obtener los posts: %v\n", err)
		return nil, err
	}
	err = cursor.All(context.Background(), &posts)
	if err != nil {
		log.Printf("Error al obtener los posts: %v\n", err)
		return nil, err
	}
	return posts, nil
}

// update post
// func UpdatePost(id string, content interface{}) error {
// 	collection := utils.GetCollection("Content")
// 	_, err := collection.UpdateOne(context.Background(), id, content)
// 	if err != nil {
// 		log.Printf("Error al actualizar el post: %v\n", err)
// 	}
// 	return err
// }
