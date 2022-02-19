package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
)

type Login interface {
	GetUser(ctx context.Context, request login.Request) (bool, error)
}
