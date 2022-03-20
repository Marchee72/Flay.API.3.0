package login

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request login.Request) (string, error)
}

type Implementation struct {
	LoginRepository providers.Login
	AuthService     providers.Authentication
}

func (useCase *Implementation) Execute(ctx context.Context, request login.Request) (string, error) {
	userID, err := useCase.LoginRepository.GetUserID(ctx, request)
	if err != nil {
		return "", err
	}
	token, err := useCase.AuthService.CreateToken(*userID)
	if err != nil {
		return "", err
	}
	return token, nil

}
