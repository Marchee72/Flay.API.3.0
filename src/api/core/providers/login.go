package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities/lw"
)

type UserRepository interface {
	GetUser(ctx context.Context, request login.Request) (*lw.UserLw, error)
	GetUserCredentials(ctx context.Context, user login.Request) (*lw.UserLw, error)
}
