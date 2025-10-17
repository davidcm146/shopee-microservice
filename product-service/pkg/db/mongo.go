package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/davidcm146/shopee-microservice/product-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	MongoClient *mongo.Client
	MongoDB     *mongo.Database
}

func InitMongo() *MongoConfig {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := config.LoadEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := config.LoadEnv("MONGO_DB", "shopee-product")

	clientOpts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		log.Fatal("MongoDB connect error: ", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping error: ", err)
	}
	fmt.Println("Connected to MongoDB")

	return &MongoConfig{
		MongoClient: client,
		MongoDB:     client.Database(dbName),
	}
}
