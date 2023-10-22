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
	filter := bson.M{"contenttype": "post"}
	cursor, err := collection.Find(context.Background(), filter)
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
func UpdatePost(id primitive.ObjectID, content domain.PostReq) error {
	collection := utils.GetCollection("Content")
	content.BeforeUpdate() // Actualiza updatedAt antes de actualizar
	filter := bson.M{"_id": id}
	update := bson.M{"$set": content}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}
