package middleware

import (
	"context"
	// "fmt"
	"net/http"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func SessionValidationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener valor de authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).SendString("No se ha enviado el token de sesión")
		}

		collection := utils.GetCollection("Session")
	
		//consultar en la base de datos
		var session bson.M
		err := collection.FindOne(context.Background(), bson.M{"sessionToken": authHeader}).Decode(&session)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString("Token de sesión inválido")
		}

		return c.Next()
	}
}
