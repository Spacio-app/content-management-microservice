package repositories

import (
	"context"

	"github.com/Spacio-app/content-management-microservice/models"
	"github.com/Spacio-app/content-management-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateContentByID(id primitive.ObjectID, updatedContent models.Courses) error {
	collection := utils.GetCollection("content")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedContent}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

// func UpdateContent(content *models.Content) error {
// 	collection := utils.GetCollection("content")
// 	filter := bson.M{"_id": content.ID}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"title":     content.Title,
// 			"content":   content.Content,
// 			"image_url": content.ImageURL,
// 			"video_url": content.VideoURL,
// 			"link_url":  content.LinkURL,
// 		},
// 	}
// 	_, err := collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Println("Error al actualizar el contenido en el repositorio:", err)
// 	}
// 	return err
// }
