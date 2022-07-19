package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuildingRepository interface {
	GetBuildingById(ctx context.Context, buildingID primitive.ObjectID) (*entities.Building, error)
	GetBuildings(ctx context.Context, ids []primitive.ObjectID) ([]entities.Building, error)
}
