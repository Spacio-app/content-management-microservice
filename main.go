package main

// "content-management-microservice/utils"
import (
	"content-management-microservice/src/utils"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	utils.InitDatabase()

	// Obtener una referencia a la colección
	collection := utils.GetCollection("<content>") // <colección>

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
}
