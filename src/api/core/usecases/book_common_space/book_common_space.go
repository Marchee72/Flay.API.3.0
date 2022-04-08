package book_common_space

import (
	"context"
	"errors"

	book_common_space "flay-api-v3.0/src/api/core/contracts/book_common_area"
)

type UseCase interface {
	Execute(ctx context.Context, request book_common_space.Request) (*book_common_space.Response, error)
}

type Implementation struct{}

func (useCase *Implementation) Execute(ctx context.Context, request book_common_space.Request) (*book_common_space.Response, error) {
	//Check if user has active penalties
	//Check if common space is abailable that day at that time
	//Book space or return error if its not abailable
	return nil, errors.New("")
}
