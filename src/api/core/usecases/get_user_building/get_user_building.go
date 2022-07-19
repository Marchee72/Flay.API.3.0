package get_user_building

import (
	"context"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, userID primitive.ObjectID, userType constants.UserType)
}

type Implementation struct {
	BuildingRepository providers.BuildingRepository
}

func (useCase Implementation) Execute(ctx context.Context, userID primitive.ObjectID, userType constants.UserType) {

}
