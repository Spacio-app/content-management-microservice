package main

// "content-management-microservice/utils"
import (
	"context"
	"fmt"
	"log"

	"github.com/Spacio-app/content-management-microservice/routes"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Inicializar la base de datos antes de configurar el servidor
	utils.InitDatabase()

	// Crear instancia de Fiber
	app := fiber.New()

	// Configurar middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Configurar rutas
	routes.SetupRoutes(app)

	// Obtener una referencia a la colección
	collection := utils.GetCollection("content") // <colección>

	// Consulta
	filter := bson.M{"campo": "valor"}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterar a través de los resultados
	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Iniciar el servidor
	port := 3000
	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	} else {
		fmt.Printf("Servidor en ejecución en el puerto %d\n", port)
	}
}
