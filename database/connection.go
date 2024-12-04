package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/roohanyh/lila_p1/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	MONGO_URI := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.Env.MONGO_USER, config.Env.MONGO_PASS, config.Env.MONGO_HOST, config.Env.MONGO_PORT)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalf("error connecting to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("error pinging MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB successfully")
	return client
}

func OpenCollection(databaseName, collectionName string) *mongo.Collection {
	client := DBInstance()

	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
