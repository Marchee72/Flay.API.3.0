package store

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BuildingRepository struct {
	Buildings *mongo.Collection
}

func (repository *BuildingRepository) GetRepository(ctx context.Context, buildingID primitive.ObjectID) (*entities.Building, error) {
	var result entities.Building
	if err := repository.Buildings.FindOne(ctx, bson.M{"_id": buildingID}).Decode(&result); err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}
