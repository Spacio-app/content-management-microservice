package repositories

import (
	"context"
	"fmt"

	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// get content by id
func GetContentByID(id primitive.ObjectID) (models.GenericContent, error) {
	collection := utils.GetCollection("Content")
	var content models.GenericContent
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&content)
	if err != nil {
		return content, err
	}
	return content, nil
}

func GetAllContent() ([]models.GenericContent, error) {
	collection := utils.GetCollection("Content")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var content []models.GenericContent
	if err := cursor.All(context.Background(), &content); err != nil {
		return nil, err
	}
	fmt.Println("content", content)
	return content, nil
}
func GetContentFeedOrderByDate(skip, limit int) ([]models.GenericContent, error) {
	collection := utils.GetCollection("Content")

	options := options.Find()
	options.SetSort(bson.M{"createdat": 1}) // Ordenar por fecha de creación

	// Aplicar Skip a la colección
	options.SetSkip(int64(skip))

	cursor, err := collection.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var content []models.GenericContent
	if err := cursor.All(context.Background(), &content); err != nil {
		return nil, err
	}

	return content, nil
}
