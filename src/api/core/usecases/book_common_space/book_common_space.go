package book_common_space

import (
	"context"
	"errors"

	book_common_space "flay-api-v3.0/src/api/core/contracts/book_common_area"
)

type UseCase interface {
	Execute(ctx *context.Context, request book_common_space.Request) error
}

type Implementation struct{}

func (useCase *Implementation) Execute(ctx *context.Context, request book_common_space.Request) error {

	return errors.New("")
}
