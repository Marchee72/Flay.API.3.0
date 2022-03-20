package store

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repository *LoginRepository) GetUserID(ctx context.Context, user login.Request) (*primitive.ObjectID, error) {
	coll := repository.DBClient.Collection("users_login")
	opt := options.FindOne().SetProjection(bson.D{{"_id", 1}})
	var result primitive.ObjectID
	err := coll.FindOne(ctx, bson.M{"username": user.Username, "password": user.Password}, opt).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}
