package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/domain/models"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCourse(content domain.CourseReq) error {
	fmt.Println(content)
	collection := utils.GetCollection("Content")
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

// get all courses from content collection
func GetAllCourses() ([]models.Courses, error) {
	collection := utils.GetCollection("Content")
	filter := bson.M{"contenttype": "course"}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var course []models.Courses
	if err := cursor.All(context.Background(), &course); err != nil {
		return nil, err
	}
	return course, nil

}

// update course
func UpdateCourse(id primitive.ObjectID, content domain.CourseReq) error {
	collection := utils.GetCollection("Content")
	content.BeforeUpdate() // Actualiza updatedAt antes de actualizar
	//mantener createdAt original
	filter := bson.M{"_id": id}
	update := bson.M{"$set": content}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}

// Initialize the Cloudinary client using environment variables
// cloudinaryURL := os.Getenv("CLOUDINARY_URL")
// cld, err := cloudinary.NewFromURL(cloudinaryURL)
// if err != nil {
// 	log.Fatalf("Error initializing Cloudinary client: %v", err)
// }
// // Upload each video link to Cloudinary
// updatedContent, err := uploadVideoLinksToCloudinary(cld, content)
// if err != nil {
// 	log.Println("Error uploading video links to Cloudinary:", err)
// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 		"error": "Error uploading video links",
// 	})
// }

// // content.VideosURL = uploadedVideoLinks
// // Upload each video link to Cloudinary
// // uploadedVideoLinks := make([]string, len(content.VideosURL))
// // for i, videoURL := range content.VideosURL {
// // 	uploadedVideoLink, err := uploadVideoLinkToCloudinary(videoURL)
// // 	if err != nil {
// // 		log.Printf("Error uploading video link %d to Cloudinary: %v", i, err)
// // 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// // 			"error": "Error uploading video link",
// // 		})
// // 	}
// // 	uploadedVideoLinks[i] = uploadedVideoLink
// // }

// // Set the Cloudinary URL of the uploaded video link in your content structure
// // content.VideosURL = uploadedVideoLinks

// Function to upload video links to Cloudinary
func uploadVideoLinksToCloudinary(cld *cloudinary.Cloudinary, content *domain.CourseReq) (*domain.CourseReq, error) {
	var updatedVideosURL []string

	for _, videoLink := range content.VideosURL {
		// Upload the video link to Cloudinary
		uploadResult, err := cld.Upload.Upload(context.TODO(), videoLink, uploader.UploadParams{})
		if err != nil {
			return nil, err
		}

		// Append the URL of the uploaded video link to the updatedVideosURL slice
		updatedVideosURL = append(updatedVideosURL, uploadResult.SecureURL)
	}

	// Update the content structure with the Cloudinary URLs
	content.VideosURL = updatedVideosURL

	return content, nil
}
