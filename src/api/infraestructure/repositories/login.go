package store

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoginRepository struct {
	DBClient *mongo.Database
}

func (repository *LoginRepository) GetUser(ctx context.Context, user login.Request) (*lw.UserLw, error) {
	coll := repository.DBClient.Collection("users_login")
	var result lw.UserLw
	err := coll.FindOne(ctx, bson.M{"username": user.Username, "password": user.Password}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}

func (repository *LoginRepository) GetUserCredentials(ctx context.Context, user login.Request) (*lw.UserLw, error) {
	coll := repository.DBClient.Collection("users_login")
	opt := options.FindOne().SetProjection(bson.D{{"_id", 1}, {"type", 1}})
	var result lw.UserLw
	err := coll.FindOne(ctx, bson.M{"username": user.Username, "password": user.Password}, opt).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}
