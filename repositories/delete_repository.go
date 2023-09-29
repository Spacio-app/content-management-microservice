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
