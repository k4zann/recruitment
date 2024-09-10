package main

import (
	"context"
	"log"
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

	var err error
	mongoClient, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI(
			"mongodb+srv://arshataitkozha:010arshat@recruitment.ak2m2.mongodb.net/?retryWrites=true&w=majority&appName=Recruitment",
		),
	)
	if err != nil || mongoClient == nil {
		log.Fatal("Error connecting to MongoDB ", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging MongoDB ", err)
	}

	log.Println("Connected to MongoDB")
}

func main() {
	// Ensure mongoClient is not nil before disconnecting
	if mongoClient == nil {
		log.Fatal("mongoClient is nil, cannot proceed")
	}
	defer mongoClient.Disconnect(context.Background())

	log.Println(mongoClient)

	// Initialize the database
	db := mongoClient.Database("recruitment")
	if db == nil {
		log.Fatal("Failed to initialize ApiServer: Database is nil")
	}

	// Start the API server
	server := api.NewApiServer(
		":8080",
		db,
	)

	if err := server.Run(); err != nil {
		log.Fatal("Error running server ", err)
	}
	log.Println("Starting server on :8080")
}
