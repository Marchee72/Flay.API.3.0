package get_file

import (
	"context"

	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, contentId primitive.ObjectID) ([]byte, error)
}

type Implementation struct {
	FileRepository providers.FileRepository
}

func (usecase Implementation) Execute(ctx context.Context, contenId primitive.ObjectID) ([]byte, error) {
	file, err := usecase.FileRepository.GetFile(ctx, contenId)
	if err != nil {
		return nil, err
	}
	return file, nil
}
