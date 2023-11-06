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
func GetContentFeedOrderByDate(skip int, limitInt int) ([]models.Feed, error) {
	collection := utils.GetCollection("Feed")

	options := options.Find()
	options.SetSort(bson.M{"createdat": -1}) // Ordenar por fecha de creación
	options.SetSkip(int64(skip))
	options.SetLimit(int64(limitInt)) // Aplicar el límite

	cursor, err := collection.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var content []models.Feed
	if err := cursor.All(context.Background(), &content); err != nil {
		return nil, err
	}

	return content, nil
}

func GetContentByAuthor(author string) ([]models.GenericContent, error) {
	collection := utils.GetCollection("Content")

	// Define un filtro para buscar contenido por autor
	filter := bson.M{"author": author}

	cursor, err := collection.Find(context.Background(), filter)
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

func GetContentByIDFeed(id primitive.ObjectID) (models.Feed, error) {
	collection := utils.GetCollection("Feed")
	var content models.Feed
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&content)
	if err != nil {
		return content, err
	}
	return content, nil
}
