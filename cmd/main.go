package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/norbux/url-shortnr/pkg/shortnr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}

	mongoUri := os.Getenv("MONGO_URI")
	//mongoDatabase := os.Getenv("MONGO_DATABASE")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	srv := shortnr.NewService(client)

	app := fiber.New()

	app.Get("/b", srv.GetURLB62)
	app.Get("/x", srv.GetURLXxh3)

	app.Listen(":8080")
}
