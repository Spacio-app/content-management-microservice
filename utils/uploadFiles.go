package utils

import (
	"github.com/gofiber/fiber/v2"

	"fmt"
	"path/filepath"

	"github.com/google/uuid"
)

// func ProcessUploadedFiles(c *fiber.Ctx, formField string) ([]string, []string, error) {
// 	// Accede a los archivos cargados
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	files := form.File[formField]

// 	filesURLS := []string{}
// 	filesIDs := []string{}

// 	for _, file := range files {
// 		// Abra el archivo
// 		src, err := file.Open()
// 		if err != nil {
// 			return nil, nil, err
// 		}
// 		defer src.Close()

// 		// Procesa el archivo y carga en Cloudinary (ajusta según tu función)
// 		publicID, secureURL, err := UploadContentToCloudinary(src)

// 		filesURLS = append(filesURLS, secureURL)
// 		filesIDs = append(filesIDs, publicID)

// 		if err != nil {
// 			return nil, nil, err
// 		}

// 	}

// 	return filesURLS, filesIDs, err
// }

func ProcessUploadedFiles(c *fiber.Ctx, formField string) ([]string, []string, error) {
	// Accede a los archivos cargados
	form, err := c.MultipartForm()
	if err != nil {
		return nil, nil, err
	}

	files := form.File[formField]

	filesURLs := []string{}
	filesPublicIDs := []string{}

	for _, file := range files {
		// Abra el archivo
		src, err := file.Open()
		if err != nil {
			return nil, nil, err
		}
		defer src.Close()

		// Obtener la extensión original del archivo
		extension := filepath.Ext(file.Filename)

		//transformar todo lo que no sea imagen o video a pdf
		if extension != ".png" && extension != ".jpg" && extension != ".jpeg" && extension != ".mp4" && extension != ".mov" && extension != ".avi" && extension != ".mkv" && extension != ".webm" {
			extension = ".pdf"
		}

		// Generar un nombre único para el archivo en Cloudinary
		publicID := generateUniqueName() + extension

		// Procesa el archivo y carga en Cloudinary
		publicID, secureURL, err := UploadContentToCloudinary(src)

		filesURLs = append(filesURLs, secureURL)
		filesPublicIDs = append(filesPublicIDs, publicID)

		if err != nil {
			return nil, nil, err
		}
	}

	return filesURLs, filesPublicIDs, nil
}

func generateUniqueName() string {
	// Generate a new UUID
	id := uuid.New()

	// Convert the UUID to a string
	idStr := id.String()

	// Output the generated UUID for verification
	fmt.Printf("UUID generated: %s\n", idStr)

	// You can add any additional logic here if needed

	// Return the UUID string as a unique name
	return idStr
}
