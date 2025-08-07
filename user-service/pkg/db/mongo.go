package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/davidcm146/shopee-microservice/user-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	MongoDB     *mongo.Database
)

func InitMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(config.LoadEnv("MONGO_URI", "mongodb://localhost:27017"))
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("MongoDB connect error: ", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping error: ", err)
	}

	MongoClient = client
	MongoDB = client.Database(config.LoadEnv("MONGO_DB", "shopee-user"))
	fmt.Println("Connected to MongoDB")
}
