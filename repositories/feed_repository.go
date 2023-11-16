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
		fmt.Println(err)
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
		CommentID: comment.CommentID,
		Comment:   comment.Comment,
		Author:    authormodel,
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

func GetFeedByAuthor(AuthorID string) ([]models.Feed, error) {
	collection := utils.GetCollection("Feed")
	log.Println("Obteniendo los posts de la base de datos...")
	filter := bson.M{"authorID": AuthorID}
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

func DeleteFeedComment(objectID primitive.ObjectID, commentID primitive.ObjectID) error {
	// Obtener el post por su ID
	existingPost, err := GetContentByIDFeed(objectID)
	if err != nil {
		return err
	}

	// Encontrar y eliminar el comentario por su ID
	var updatedComments []models.FeedComments
	for _, comment := range existingPost.Comments {
		if comment.CommentID != commentID {
			updatedComments = append(updatedComments, comment)
		}
	}
	updatedPost := models.Feed{
		Comments: updatedComments,
	}
	// Actualizar el post en la base de datos
	err = UpdateFeedComments(objectID, updatedPost)
	if err != nil {
		return err
	}

	return nil
}
