package utils

import (
	"log"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ProcessUploadedFiles(c *fiber.Ctx, file *multipart.FileHeader, isVideo bool, miniature *multipart.FileHeader) (string, string, string, string, error) {

	// Abra el archivo
	src, err := file.Open()
	if err != nil {
		return "", "", "", err
	}
	defer src.Close()

	//abrir miniatura
	srcMiniature, err := miniature.Open()
	if err != nil {
		miniatureURL := ""
	}
	defer srcMiniature.Close()

	if srcMiniature != nil {
		// Procesa el archivo y carga en Cloudinary
		PublicIDMiniature, miniatureURL, err := UploadContentToCloudinary(srcMiniature)
		if err != nil {
			miniatureURL := ""
		}
	}
	else {
		miniatureURL := ""
	}



	// Procesa el archivo y carga en Cloudinary
	publicID, secureURL, err := UploadContentToCloudinary(src)

	if err != nil {
		return "", "", "", err
	}

	//preguntar si miniatureURL es nulo
	if miniatureURL != nil {
		return secureURL, publicID, miniatureURL, PublicIDMiniature, nil
	}
	
	log.Printf("isVideo: %v\n", isVideo)
	if isVideo {
		miniatureURL := ""
		firstVideoURL := secureURL
		// log.Println("firstVideoURL", firstVideoURL)
		// fmt.Println("firstVideoURL", firstVideoURL)

		// Lista de formatos de video admitidos
		videoFormats := []string{".mp4", ".mov", ".webm", ".avi", ".mkv", ".wmv", ".flv", ".3gp", ".mpeg", ".mpg", ".m4v"}

		// Itera a través de los formatos de video y cambia la extensión a .jpg
		for _, format := range videoFormats {
			firstVideoURL = strings.Replace(firstVideoURL, format, ".jpg", 1)
		}

		// Almacena la URL de la miniatura
		miniatureURL = firstVideoURL
		log.Println("miniatureURL", miniatureURL)
	}

	return secureURL, publicID, miniatureURL, nil
}

