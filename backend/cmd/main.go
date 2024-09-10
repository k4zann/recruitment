package main

import (
	"context"
	"log"
	"os"
	"recruitment/cmd/api"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	mongoClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(
			os.Getenv("DB_LINK"),
		),
	)
	if err != nil {
		log.Fatal("Error connecting to MongoDB ", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging MongoDB ", err)
	}

	log.Println("Connected to MongoDB")
}

func main() {

	defer mongoClient.Disconnect(context.Background())

	log.Println(mongoClient)

	db := mongoClient.Database("recruitment")
	if db == nil {
		log.Fatal("Failed to initialize ApiServer: Database is nil")
	}

	server := api.NewApiServer(
		":8080",
		db,
	)

	if err := server.Run(); err != nil {
		log.Fatal("Error running server ", err)
	}
	log.Println("Starting server on :8080")

}
