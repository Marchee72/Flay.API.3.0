package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	GetUser(ctx context.Context, request login.Request) (*lw.UserLw, error)
	GetUserById(ctx context.Context, userID primitive.ObjectID) (*entities.User, error)
	GetUserCredentials(ctx context.Context, user login.Request) (*entities.Credentials, error)
}
