package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_URI = "mongodb://localhost:27017"
)

func Connect() (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		return nil, err
	}
	db := client.Database("Flay")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
