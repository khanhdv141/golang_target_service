package repository

import (
	"CMS/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoClient *mongo.Client

func InitMongoClient() (*mongo.Client, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	uri := fmt.Sprintf("mongodb://%s:%d",
		config.ApplicationConfig.Mongo.Host,
		config.ApplicationConfig.Mongo.Port,
	)
	clientOptions := options.Client().ApplyURI(uri)
	MongoClient, err := mongo.Connect(ctx, clientOptions)
	return MongoClient, cancel, err
}
