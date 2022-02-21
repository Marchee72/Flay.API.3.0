package login

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request login.Request) (*entities.UserLogin, error)
}

type Implementation struct {
	LoginRepository providers.Login
}

func (useCase *Implementation) Execute(ctx context.Context, request login.Request) (*entities.UserLogin, error) {
	successfulLogin, err := useCase.LoginRepository.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return successfulLogin, nil

}
