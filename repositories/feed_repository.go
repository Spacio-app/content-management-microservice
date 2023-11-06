package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFeed(content domain.FeedReq) error {
	fmt.Println(content)
	collection := utils.GetCollection("Feed")
	content.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar
	log.Printf("Content: %v\n", content)
	log.Println("Insertando datos en la base de datos...")

	_, err := collection.InsertOne(context.Background(), content)
	if err != nil {
		log.Printf("Error al crear el curso: %v\n", err)
		return err
	}
	return nil
}
func UpdatePostComments(objectID primitive.ObjectID, updatedComments []domain.FeedCommentsReq) error {
	// Buscar el post por su ID
	existingPost, err := GetContentByID(objectID)
	if err != nil {
		return err
	}

	// Iterar sobre los comentarios actualizados y agregarlos al post existente
	for _, comment := range updatedComments {
		commentModel := models.FeedComments{
			Comment:     comment.Comment,
			Author:      comment.Author.Name,
			AuthorPhoto: comment.Author.Photo,
		}
		commentModel.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar
		existingPost.FeedComments = append(existingPost.FeedComments, commentModel)
	}

	// Actualizar el post en la base de datos
	err = UpdateFeedComments(objectID, existingPost)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFeedComments(objectID primitive.ObjectID, updatedPost models.GenericContent) error {
	// Actualizar el post en la base de datos
	collection := utils.GetCollection("Feed")
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": objectID}, bson.M{"$set": updatedPost})
	if err != nil {
		return err
	}

	return nil
}

func GetAllFeeds() ([]models.Feed, error) {
	collection := utils.GetCollection("Feed")
	log.Println("Obteniendo los posts de la base de datos...")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error al obtener los posts: %v\n", err)
		return nil, err
	}
	log.Println("cursor", cursor)
	var feeds []models.Feed
	err = cursor.All(context.Background(), &feeds)
	if err != nil {
		log.Printf("Error al obtener los posts: %v\n", err)
		return nil, err
	}
	log.Println("Posts obtenidos correctamente", feeds)
	return feeds, nil
}

// func generateFeedFromContent(content models.Feed) error{
// 	collection := utils.GetCollection("Feed")
// 	content.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar

// 	//crear estructura feed para insertar en la base de datos
// 	feed := models.Feed{
// 		Title:       content.Title,
// 		Description: content.Description,
// 		ContentType: content.ContentType,
// 		ContentID:   content.ID,
// 		AuthorID:    content.AuthorID,
// 		Author:	  content.Author,
// 		AuthorPhoto: content.AuthorPhoto,
// 	}
// 	_, err := collection.InsertOne(context.Background(), feed)
// 	if err != nil {
// 		log.Printf("Error al crear el curso: %v\n", err)
// 	}
// 	return err
// }
