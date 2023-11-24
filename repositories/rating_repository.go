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

func HasRated(rating domain.RatingReq, contentID primitive.ObjectID) (bool, error) {
	collection := utils.GetCollection("Ratings")
	filter := bson.M{"content_id": contentID, "author.email": rating.Author.Email}
	var result models.Rating
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return false, nil
	}
	return true, nil
}
func RateContent(rating domain.RatingReq) error {
	collection := utils.GetCollection("Ratings")
	rating.BeforeInsert()
	_, err := collection.InsertOne(context.Background(), rating)
	if err != nil {
		log.Printf("Error al crear el rating: %v\n", err)
	}
	return err
}
func GetRatingCount(contentID primitive.ObjectID) (int64, error) {
	collection := utils.GetCollection("Ratings")
	filter := bson.M{"content_id": contentID}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func GetRatingAverage(contentID primitive.ObjectID) (float64, error) {
	collection := utils.GetCollection("Ratings")
	filter := bson.M{"content_id": contentID}
	var result []models.Rating
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return 0, err
	}
	var total float64
	for _, rating := range result {
		total += rating.Rating
	}
	return total / float64(len(result)), nil
}
