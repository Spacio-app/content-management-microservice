package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

// Inicializar Cloudinary
func InitCloudinary() (*cloudinary.Cloudinary, error) {
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error cargando variables de entorno: %v", err)
	}

	// Obtener credenciales de Cloudinary desde variables de entorno
	apiKey := "564161591792583"
	apiSecret := "En_yFQLYIpFpDHqKsWTn0u5L1C0"
	cloudName := "ddmhgap5x"
	log.Println("apiSecret", apiSecret)
	// Crear una instancia de Cloudinary
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, fmt.Errorf("Error inicializando el cliente de Cloudinary: %v", err)
	}

	return cld, nil
}

// UploadVideoToCloudinary sube un video a Cloudinary y devuelve su publicID y URL segura.
func UploadContentToCloudinary(file interface{}) (string, string, error) {
	// Inicializar el cliente de Cloudinary
	cld, err := InitCloudinary()
	if err != nil {
		return "", "", fmt.Errorf("Error inicializando el cliente de Cloudinary: %v", err)
	}
	log.Println("Contenido subido", file)
	// Subir el video a Cloudinary
	uploadResult, err := cld.Upload.Upload(context.TODO(), file, uploader.UploadParams{})
	if err != nil {
		return "", "", fmt.Errorf("Error al subir el video a Cloudinary: %v", err)
	}

	// Obtener el publicID y la URL segura del video
	publicID := uploadResult.PublicID
	secureURL := uploadResult.SecureURL

	return publicID, secureURL, nil
}

// Obtener detalles de un recurso en Cloudinary por su PublicID
func GetCloudinaryResourceDetails(publicID string) (*admin.AssetResult, error) {
	cld, err := InitCloudinary()
	if err != nil {
		return nil, err
	}

	// Obtener detalles del recurso por su PublicID
	resp, err := cld.Admin.Asset(context.TODO(), admin.AssetParams{PublicID: publicID})
	if err != nil {
		return nil, fmt.Errorf("Error al obtener detalles del recurso en Cloudinary: %v", err)
	}

	return resp, nil
}

// Obtener la URL segura de un recurso en Cloudinary por su PublicID
func GetCloudinaryResourceURL(publicID string) (string, error) {
	cld, err := InitCloudinary()
	if err != nil {
		return "", err
	}

	// Obtener detalles del recurso por su PublicID
	resp, err := cld.Admin.Asset(context.TODO(), admin.AssetParams{PublicID: publicID})
	if err != nil {
		return "", fmt.Errorf("Error al obtener detalles del recurso en Cloudinary: %v", err)
	}

	return resp.SecureURL, nil
}

// package utils

// // Import the required packages for upload and admin.

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"fmt"

// 	"github.com/cloudinary/cloudinary-go/v2"
// 	"github.com/cloudinary/cloudinary-go/v2/api"
// 	"github.com/cloudinary/cloudinary-go/v2/api/admin"
// 	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
// 	"github.com/joho/godotenv"
// )

// func InitCloudinary() {
// 	// cld, _ := cloudinary.New()
// 	// Initialize the Cloudinary client using environment variables
// 	// cloudinaryURL := os.Getenv("CLOUDINARY_URL")
// 	apiKey := os.Getenv("CLOUDINARY_API_KEY")
// 	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
// 	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
// 	cld, _ := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
// 	if err != nil {
// 		log.Fatalf("Error initializing Cloudinary client: %v", err)
// 	}

// 	// Upload the my_picture.jpg image and set the PublicID to "my_image".

// 	resp, err := cld.Upload.Upload(ctx, "my_picture.jpg", uploader.UploadParams{PublicID: "my_image"});
// 	if err != nil {
// 		log.Fatalf("Failed to upload image: %v", err)
// 	}
// 	// Get details about the image with PublicID "my_image" and log the secure URL.

// 	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "my_image"});
// 	if err != nil {
// 	log.Println(resp.SecureURL)
// 	}
// 	// Instantiate an object for the asset with public ID "my_image"
// 	my_image, err := cld.Image("my_image")
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// 	// Add the transformation
// 	my_image.Transformation = "c_fill,h_250,w_250"

// 	// Generate and print the delivery URL
// 	url, err := my_image.String()
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// }
