package book_common_space

import (
	"context"
	"errors"

	book_common_space "flay-api-v3.0/src/api/core/contracts/book_common_space"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request book_common_space.Request) (*book_common_space.Response, error)
}

type Implementation struct {
	PenaltyRepository providers.PenaltyRepository
}

func (useCase *Implementation) Execute(ctx context.Context, request book_common_space.Request) (*book_common_space.Response, error) {
	//Check if user has active penalties
	penalty, err := useCase.PenaltyRepository.GetActivePenalties(ctx, request.UserID)
	if err != nil {
		return nil, err
	}
	response := book_common_space.Response{}
	if penalty != nil {
		response.SetPenalty(*penalty)
		return &response, nil
	}
	//Check if common space is abailable that day at that timeSs
	//Book space or return error if its not abailable
	return nil, errors.New("")
}
