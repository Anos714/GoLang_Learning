package db

import (
	"context"
	"fmt"
	"notes-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {

	// Set a 10-second timeout context for the initial connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configure client options
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	// Connect to Mongo
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("Mongo connection failed: %w",err)
	}


	// ping the db
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Mongo ping failed: %w", err)
	}

	fmt.Println("Connected successfully to MongoDB!")

	db := client.Database(cfg.MongoDB)
	return client, db, nil
}


func Disconnect(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("Mongo disconnect failed: %w", err)
	}

	fmt.Println("Disconnected from MongoDB.")
	return nil
}
