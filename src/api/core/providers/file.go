package providers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileRepository interface {
	SaveFile(ctx context.Context, file []byte, contentId primitive.ObjectID) error
	GetFile(ctx context.Context, contentId primitive.ObjectID) ([]byte, error)
}
