package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirhmvfx/react-golang-blog/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func initMongoDb() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load env", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connected")
	return client

}

func main() {
	fmt.Println("Blog Api")

	app := fiber.New()

	mongoClient := initMongoDb()
	router.Router(app, mongoClient)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}

	log.Fatal(app.Listen(":" + PORT))
}
