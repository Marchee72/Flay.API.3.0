package save_announcement

import (
	"context"
	"errors"

	"flay-api-v3.0/src/api/core/contracts/save_announcement"
	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, request save_announcement.Request, userID primitive.ObjectID) error
}
type Implementation struct {
	AnnouncementRepository providers.AnnouncementRepository
	UserRepository         providers.UserRepository
}

func (usecase Implementation) Execute(ctx context.Context, request save_announcement.Request, userID primitive.ObjectID) error {
	return errors.New("")
}
