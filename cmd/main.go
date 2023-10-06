package main

// "content-management-microservice/utils"
import (
	"context"
	"fmt"
	"log"
	"os"

	//	"github.com/joho/godotenv"

	"github.com/Spacio-app/content-management-microservice/routes"
	"github.com/Spacio-app/content-management-microservice/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

// func configureCloudinary() *cloudinary.Cloudinary {
// 	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
// 	cld, err := cloudinary.NewFromURL(cloudinaryURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return cld
// }

func main() {
	// Inicializar la base de datos antes de configurar el servidor
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	utils.InitCloudinary()
	utils.InitDatabase()
	// Crear instancia de Fiber
	app := fiber.New()

	// Configurar middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	// app.Use(middleware.SessionValidationMiddleware())

	// Configurar rutas
	// Configurar el archivo de registro
	logFile, err := os.Create("app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	// Configurar el registro para que escriba en el archivo
	log.SetOutput(logFile)

	routes.SetupRoutes(app)

	// Obtener una referencia a la colección
	collection := utils.GetCollection("Content") // <colección>

	// Consulta
	// filter := bson.M{"campo": "valor"}
	filter := bson.M{} // <filtro>
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" // Valor predeterminado en caso de que la variable de entorno no esté configurada
	}
	err = app.Listen(":" + port) // Agrega ":" antes del número del puerto
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
		return
	} else {
		fmt.Printf("Servidor en ejecución en el puerto %s\n", port)
	}
}
