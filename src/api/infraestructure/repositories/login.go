package store

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRepository struct {
	DBClient *mongo.Database
}

func (repository *LoginRepository) GetUser(ctx context.Context, user login.Request) (bool, error) {
	coll := repository.DBClient.Collection("users")
	var result bson.M
	err := coll.FindOne(ctx, bson.D{{"username", user.Username}, {"password", user.Password}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, err
	}
	return true, nil
}
