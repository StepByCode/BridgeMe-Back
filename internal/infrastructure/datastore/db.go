package datastore

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDB(connectionString string) (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, nil, err
	}

	log.Println("Connected to MongoDB!")
	db := client.Database("bridgeme")

	return client, db, nil
}

func CloseDB(client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}

	log.Println("Connection to MongoDB closed.")
}
