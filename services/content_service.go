// services/content_service.go
package services

import (
	"log"

	"github.com/Spacio-app/content-management-microservice/models"
	"github.com/Spacio-app/content-management-microservice/repositories"
)

func CreateContent(content *models.Content) error {
	if err := repositories.CreateContent(content); err != nil {
		log.Println("Error en el servicio al crear el contenido:", err)
		return err
	}
	return nil
}

func GetContentByID(id string) (*models.Content, error) {
	content, err := repositories.GetContentByID(id)
	if err != nil {
		log.Println("Error en el servicio al obtener el contenido por ID:", err)
		return nil, err
	}
	return content, nil
}
func DeleteContentByID(id string) error {
	if err := repositories.DeleteContentByID(id); err != nil {
		log.Println("Error en el servicio al eliminar el contenido:", err)
		return err
	}
	return nil
}

func UpdateContent(content *models.Content) error {
	if err := repositories.UpdateContent(content); err != nil {
		log.Println("Error en el servicio al actualizar el contenido:", err)
		return err
	}
	return nil
}

// Continuar con las funciones para actualizar y eliminar contenido...
