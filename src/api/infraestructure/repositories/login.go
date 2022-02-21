package store

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoginRepository struct {
	DBClient *mongo.Database
}

func (repository *LoginRepository) GetUser(ctx context.Context, user login.Request) (*entities.UserLogin, error) {
	coll := repository.DBClient.Collection("users_login")
	var result entities.UserLogin
	err := coll.FindOne(ctx, bson.M{"username": user.Username, "password": user.Password}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}
