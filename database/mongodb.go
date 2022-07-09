package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbRepository struct {
	DB *mongo.Database
}

func NewMongoRepository(url, name string) (*MongodbRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established")
	return &MongodbRepository{DB: client.Database(name)}, nil
}
