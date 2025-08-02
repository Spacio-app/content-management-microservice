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

		collection := utils.GetCollection("Account")

		//consultar en la base de datos
		var session bson.M
		err := collection.FindOne(context.Background(), bson.M{"id_token": authHeader}).Decode(&session)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString("Token de sesión inválido")
		}
		// Decodificar el token y obtener los datos del usuario
		//  userID, userPhoto, username, err := utils.DecodeSessionToken(authHeader)
		//  if err != nil {
		// 	 return c.Status(http.StatusUnauthorized).SendString("Token de sesión inválido")
		//  }

		// Almacenar los datos del usuario en el contexto de Fiber
		//  c.Locals("userID", userID)
		//  c.Locals("userPhoto", userPhoto)
		//  c.Locals("userName", username)

		return c.Next()
	}
}
