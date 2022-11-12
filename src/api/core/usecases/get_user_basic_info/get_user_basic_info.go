package get_user_basic_info

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_user_basic_info"
	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UseCase interface {
	Execute(ctx context.Context, userID primitive.ObjectID) (*get_user_basic_info.Response, error)
}

type Implementation struct {
	UserRepository      providers.UserRepository
	BuildingRepository  providers.BuildingRepository
	ApartmentRepository providers.ApartmentRepository
}

func (usecase Implementation) Execute(ctx context.Context, userID primitive.ObjectID) (*get_user_basic_info.Response, error) {
	var response get_user_basic_info.Response
	user, err := usecase.UserRepository.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	building, err := usecase.BuildingRepository.GetBuildingById(ctx, user.Building.ID)
	if err != nil {
		return nil, err
	}
	apartment, err := usecase.ApartmentRepository.GetApartmentByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	response.NewResponse(user.ToLw(), *building, *apartment)
	return &response, nil
}
