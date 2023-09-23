package repositories

import (
	"context"

	"github.com/Spacio-app/content-management-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteContentByID(id primitive.ObjectID) error {
	collection := utils.GetCollection("Content")
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.Background(), filter)
	return err
}

// func DeleteContentByID(id string) error {
// 	collection := utils.GetCollection("content")
// 	filter := bson.M{"_id": id}
// 	_, err := collection.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		log.Println("Error al eliminar el contenido en el repositorio:", err)
// 	}
// 	return err
// }
