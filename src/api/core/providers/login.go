package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login interface {
	GetUser(ctx context.Context, request login.Request) (*entities.UserLogin, error)
	GetUserID(ctx context.Context, user login.Request) (*primitive.ObjectID, error)
}
