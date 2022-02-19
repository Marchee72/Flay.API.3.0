package login

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/login"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request login.Request) (bool, error)
}

type Implementation struct {
	LoginRepository providers.Login
}

func (useCase *Implementation) Execute(ctx context.Context, request login.Request) (bool, error) {
	successfulLogin := false
	var err error
	successfulLogin, err = useCase.LoginRepository.GetUser(ctx, request)
	if err != nil {
		return successfulLogin, err
	}
	return successfulLogin, nil

}
