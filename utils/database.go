// utils/database.go
package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func InitDatabase() {
	// Cadena de conexión de MongoDB Atlas
	connectionString := "mongodb+srv://root:root@spacio.ens4lgr.mongodb.net/?retryWrites=true&w=majority"

	// Opciones de configuración
	clientOptions := options.Client().ApplyURI(connectionString)

	// Conectar a MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Obtener una referencia a la base de datos y la colección
	db = client.Database("<spacio>")
}

// GetCollection devuelve una referencia a la colección especificada
func GetCollection(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
