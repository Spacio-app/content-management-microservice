// repositories/content_repository.go
package repositories

import (
	"context"
	"log"

	"github.com/Spacio-app/content-management-microservice/models"
	"github.com/Spacio-app/content-management-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateContent(content *models.Content) error {
	collection := utils.GetCollection("content")
	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Println("Error al crear el content:", err)
	}
	return err
}

func GetContentByID(id string) (*models.Content, error) {
	collection := utils.GetCollection("content")
	filter := bson.M{"_id": id}
	var content models.Content
	err := collection.FindOne(context.Background(), filter).Decode(&content)
	if err != nil {
		log.Println("Error al obtener el content por ID:", err)
	}
	return &content, err
}
func UpdateContent(content *models.Content) error {
	collection := utils.GetCollection("content")
	filter := bson.M{"_id": content.ID}
	update := bson.M{
		"$set": bson.M{
			"title":     content.Title,
			"content":   content.Content,
			"image_url": content.ImageURL,
			"video_url": content.VideoURL,
			"link_url":  content.LinkURL,
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Error al actualizar el contenido en el repositorio:", err)
	}
	return err
}

func DeleteContentByID(id string) error {
	collection := utils.GetCollection("content")
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error al eliminar el contenido en el repositorio:", err)
	}
	return err
}
