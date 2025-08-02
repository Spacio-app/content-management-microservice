package repositories

import (
	"context"
	"errors"
	"fmt"
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
	var ratings []models.Rating
	collection := utils.GetCollection("Ratings")
	filter := bson.M{"content_id": contentID}
	fmt.Println(filter)
	//FIND ALL RATINGS FOR CONTENT
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("Error al obtener los ratrings:", err)
		return 0, err
	}

	fmt.Println("CURSOR", cursor)

	err = cursor.All(context.Background(), &ratings)
	if err != nil {
		log.Printf("Error al obtener los ratrings: %v\n", err)
		return 0, err
	}

	//CALCULATE AVERAGE
	var sum float64
	var count int64
	//recorrer ratings y sumarlos
	fmt.Println("RATINGS", ratings)

	for _, rating := range ratings {
		//transformar rating a float64
		sum += float64(rating.Rating)
		count++
	}

	if count == 0 {
		return 0, errors.New("No hay calificaciones válidas para calcular el promedio")
	}

	average := sum / float64(count)

	fmt.Println(average)
	return average, nil
}
