package book_common_area

import (
	"context"
	"errors"
)

type UseCase interface {
	Execute(ctx *context.Context) error
}

type Implementation struct{}

func (useCase *Implementation) Execute(ctx *context.Context) error {
	return errors.New("hasta aca llegamos.")
}
