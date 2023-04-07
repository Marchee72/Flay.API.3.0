package store

import "flay-api-v3.0/vendor/go.mongodb.org/mongo-driver/mongo"

type ExpenseBucket struct {
	Expenses *mongo.Database
}
