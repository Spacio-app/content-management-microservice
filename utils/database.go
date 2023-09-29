package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db         *mongo.Database
	dbInitOnce sync.Once
)

func InitDatabase() {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error cargando el archivo .env")
		log.Fatal(err)
	}

	// Obtener la cadena de conexión de MongoDB desde las variables de entorno
	connectionString := os.Getenv("MONGODB_URI")

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
