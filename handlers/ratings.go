package handlers

import (
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/domain"
	"github.com/Spacio-app/content-management-microservice/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RateContent(c *fiber.Ctx) error {
	rating := domain.RatingReq{}
	contentID := c.Params("contentID")
	// Parsear el cuerpo JSON
	if err := c.BodyParser(&rating); err != nil {
		log.Println("Error al analizar el cuerpo de la solicitud:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error al analizar el cuerpo JSON",
		})
	}
	user, error := getUserHeader(c)
	if error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al obtener el usuario",
		})
	}
	author := domain.AuthorReq{
		Name:  user.Name,
		Photo: user.Image,
		Email: user.Email,
	}

	rating.Author = author

	//transformar contentID a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	//has rated
	hasRated, err := services.HasRated(rating, objectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al verificar si el usuario ha calificado el contenido",
		})
	}
	if hasRated {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Ya has calificado este contenido",
		})
	}
	// Enviar a servicio
	if err := services.RateContent(rating); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al calificar el contenido",
		})
	}
	return c.JSON(rating)
}

// get rating count
func GetRatingCount(c *fiber.Ctx) error {
	contentID := c.Params("contentID")
	//transformar contentID a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	count, err := services.GetRatingCount(objectID)
	if err != nil {
		fmt.Println(err)
		// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 	"error": "Error al obtener el conteo de calificaciones",
		// })
	}
	return c.JSON(fiber.Map{
		"count": count,
	})
}
func GetRatingAverage(c *fiber.Ctx) error {
	contentID := c.Params("contentID")
	//transformar contentID a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(contentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	average, err := services.GetRatingAverage(objectID)
	if err != nil {
		fmt.Println(err)
		// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 	"error": "Error al obtener el promedio de calificaciones",
		// })
	}
	return c.JSON(fiber.Map{
		"average": average,
	})
}
