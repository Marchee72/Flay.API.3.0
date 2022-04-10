package login

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/providers"
	"flay-api-v3.0/src/api/infraestructure/authentication"
)

type UseCase interface {
	Execute(ctx context.Context, request login.Request) (string, error)
}

type Implementation struct {
	UserRepository providers.UserRepository
}

func (useCase *Implementation) Execute(ctx context.Context, request login.Request) (string, error) {
	user, err := useCase.UserRepository.GetUserCredentials(ctx, request)
	if err != nil {
		return "", err
	}
	token, err := authentication.CreateToken(*user)
	if err != nil {
		return "", err
	}
	return token, nil

}
