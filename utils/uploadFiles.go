package utils

import (
	"log"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ProcessUploadedFiles(c *fiber.Ctx, file *multipart.FileHeader, isVideo bool, miniature *multipart.FileHeader) (string, string, string, string, error) {
	// Verifica si `file` no es nulo
	if file != nil {
		// Abre el archivo
		src, err := file.Open()
		if err != nil {
			return "", "", "", "", err
		}
		defer src.Close()

		// Procesa el archivo y carga en Cloudinary
		publicID, secureURL, err := UploadContentToCloudinary(src)
		if err != nil {
			return "", "", "", "", err
		}

		// Verifica si `miniature` no es nulo
		if miniature != nil {
			// Abre la miniatura
			srcMiniature, err := miniature.Open()
			if err != nil {
				log.Printf("No se proporcionó una miniatura")
			}
			defer srcMiniature.Close()

			PublicIDMiniature, miniatureURL, err := UploadContentToCloudinary(srcMiniature)
			if err != nil {
				miniatureURL = ""
			}

			// Pregunta si `miniatureURL` no es nulo
			if miniatureURL != "" {
				return secureURL, publicID, miniatureURL, PublicIDMiniature, nil
			}
		}

		log.Printf("isVideo: %v\n", isVideo)
		if isVideo {
			miniatureURL := ""
			firstVideoURL := secureURL
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

		return secureURL, publicID, "", "", nil
	}

	// Retorna valores predeterminados si `file` es nulo
	return "", "", "", "", nil
}
