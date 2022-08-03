package store

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Users *mongo.Collection
}

func (repository *UserRepository) GetUser(ctx context.Context, user login.Request) (*lw.UserLw, error) {
	var result lw.UserLw
	err := repository.Users.FindOne(ctx, bson.M{"username": user.Username, "password": user.Password}).Decode(&result)
	if err != mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}

func (repository *UserRepository) GetUserCredentials(ctx context.Context, user login.Request) (*entities.Credentials, error) {
	opt := options.FindOne().SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "username", Value: 1}, {Key: "type", Value: 1}})
	var result entities.Credentials
	err := repository.Users.FindOne(ctx, bson.M{"username": user.Username, "password": user.Password}, opt).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repository *UserRepository) GetUserById(ctx context.Context, userID primitive.ObjectID) (*entities.User, error) {
	var result entities.User
	err := repository.Users.FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
