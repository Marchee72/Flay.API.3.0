package get_buildings

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/common"
	get_Building "flay-api-v3.0/src/api/core/contracts/get_building"
	"flay-api-v3.0/src/api/core/providers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetBuildingUseCase interface {
	Execute(ctx context.Context, userId primitive.ObjectID) (*get_Building.Response, error)
}

type Implementation struct {
	BuildingRepository providers.BuildingRepository
}

func (usecase Implementation) Execute(ctx context.Context, userId primitive.ObjectID) (*get_Building.Response, error) {
	buildings, err := usecase.BuildingRepository.GetBuildingsByAdmin(ctx, userId)
	if err != nil {
		return nil, err
	}
	response := get_Building.Response{}
	for _, b := range buildings {
		response.Buildings = append(response.Buildings, common.NewBuilding(b))
	}
	return &response, nil
}
