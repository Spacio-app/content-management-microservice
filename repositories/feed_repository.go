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
func UpdatePostComments(objectID primitive.ObjectID, comment domain.FeedCommentsReq) error {
	// Buscar el post por su ID
	existingPost, err := GetContentByIDFeed(objectID)
	if err != nil {
		return err
	}
	log.Printf("existingPost: %v\n", existingPost)
	// Convertir los comentarios actualizados a modelos de comentarios
	var updatedCommentsModel []models.FeedComments

	//author model
	authormodel := models.Author{
		Name:  comment.Author.Name,
		Photo: comment.Author.Photo,
	}

	commentModel := models.FeedComments{
		Comment: comment.Comment,
		Author:  authormodel,
	}
	commentModel.BeforeInsert() // Actualiza createdAt y updatedAt antes de insertar

	log.Printf("updatedCommentsModel: %v\n", updatedCommentsModel)
	log.Printf("existingPost: %v\n", existingPost)
	log.Printf("por aqui pase")
	log.Println("updatedCommentsModel", commentModel)
	fmt.Println("existingPost", existingPost)

	// Actualizar los comentarios del post
	// existingPost.BeforeUpdate() // Actualiza createdAt y updatedAt antes de insertar
	existingPost.Comments = append(existingPost.Comments, commentModel)

	// Actualizar el post en la base de datos
	err = UpdateFeedComments(objectID, existingPost)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func UpdateFeedComments(objectID primitive.ObjectID, updatedPost models.Feed) error {
	// Actualizar el post en la base de datos
	collection := utils.GetCollection("Feed")
	filter := bson.M{"_id": objectID}
	_, err := collection.UpdateOne(context.Background(), filter,
		bson.M{"$set": bson.M{
			"comments": updatedPost.Comments,
		}})
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
