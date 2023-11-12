package utils

import (
	"log"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ProcessUploadedFiles(c *fiber.Ctx, file *multipart.FileHeader, isVideo bool, miniature *multipart.FileHeader) (string, string, string, string, error) {
	var secureURL, publicID, miniatureURL, PublicIDMiniature string

	if file != nil {
		// Procesa el archivo y carga en Cloudinary
		src, err := file.Open()
		if err != nil {
			return "", "", "", "", err
		}
		defer src.Close()

		publicID, secureURL, err = UploadContentToCloudinary(src)
		if err != nil {
			return "", "", "", "", err
		}
	}

	if miniature != nil {
		// Procesa la miniatura y carga en Cloudinary
		srcMiniature, err := miniature.Open()
		if err != nil {
			log.Printf("No se proporcionó una miniatura")
		} else {
			defer srcMiniature.Close()
			PublicIDMiniature, miniatureURL, err = UploadContentToCloudinary(srcMiniature)
			if err != nil {
				miniatureURL = ""
			}
		}
	}

	if isVideo {
		miniatureURL := ""
		firstVideoURL := secureURL
		// Lista de formatos de video admitidos
		videoFormats := []string{".mp4", ".mov", ".webm", ".avi", ".mkv", ".wmv", ".flv", ".3gp", ".mpeg", ".mpg", ".m4v"}
		// Itera a través de los formatos de video y cambia la extensión a .jpg
		for _, format := range videoFormats {
			firstVideoURL = strings.Replace(firstVideoURL, format, ".png", 1)
		}
		// Almacena la URL de la miniatura
		miniatureURL = firstVideoURL
		return secureURL, publicID, miniatureURL, PublicIDMiniature, nil
	}

	return secureURL, publicID, miniatureURL, PublicIDMiniature, nil
}
