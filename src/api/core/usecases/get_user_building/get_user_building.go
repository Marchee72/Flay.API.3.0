package get_user_building

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_user_building"
	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, userID primitive.ObjectID) (*get_user_building.Response, error)
}

type Implementation struct {
	BuildingRepository providers.BuildingRepository
	UserRepository     providers.UserRepository
}

func (useCase Implementation) Execute(ctx context.Context, userID primitive.ObjectID) (*get_user_building.Response, error) {
	user, err := useCase.UserRepository.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	building, err := useCase.BuildingRepository.GetBuildingById(ctx, user.Building.ID)
	if err != nil {
		return nil, err
	}
	response := get_user_building.Response{
		ID:   building.ID,
		Name: building.Name,
	}
	response.SetCommonSpaces(building.CommonSpaces)
	return &response, nil
}
