package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities"
)

type Login interface {
	GetUser(ctx context.Context, request login.Request) (*entities.UserLogin, error)
}
