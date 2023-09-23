package utils

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//	"github.com/joho/godotenv"
	"fmt"
)

var (
	db         *mongo.Database
	dbInitOnce sync.Once
)

func InitDatabase() {
	// Cadena de conexión de MongoDB Atlas
	// if err := godotenv.Load(); err != nil {
	//     fmt.Println("Error cargando el archivo .env")
	//     os.Exit(1)
	// }
	//connectionString := os.Getenv("MONGODB_URI")
	connectionString := "mongodb+srv://spacio:yrhxtXKZd5O1EUVa@spacio.b0dcbnn.mongodb.net/?retryWrites=true&w=majority"

	fmt.Printf("connectionString: %s \n", connectionString)
	// Opciones de configuración
	clientOptions := options.Client().ApplyURI(connectionString)

	// Conectar a MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Asignar la referencia a la base de datos
	db = client.Database("spacio")
}

// GetCollection devuelve una referencia a la colección especificada
func GetCollection(collectionName string) *mongo.Collection {
	dbInitOnce.Do(func() {
		InitDatabase() // Asegurarse de que la conexión esté inicializada
	})

	return db.Collection(collectionName)
}
