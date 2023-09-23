package repositories

import (
	"context"

	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetContentByID(id primitive.ObjectID) (*models.Courses, error) {
	collection := utils.GetCollection("Content")
	filter := bson.M{"_id": id}

	var content models.Courses
	err := collection.FindOne(context.Background(), filter).Decode(&content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}
func GetAllContent() ([]models.Courses, error) {
	collection := utils.GetCollection("Content")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var content []models.Courses
	if err := cursor.All(context.Background(), &content); err != nil {
		return nil, err
	}

	return content, nil
}
