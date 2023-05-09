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

func (repository *BuildingRepository) GetBuildingById(ctx context.Context, buildingID primitive.ObjectID) (*entities.Building, error) {
	var result entities.Building
	if err := repository.Buildings.FindOne(ctx, bson.M{"_id": buildingID}).Decode(&result); err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}

func (repository *BuildingRepository) GetBuildings(ctx context.Context, ids []primitive.ObjectID) ([]entities.Building, error) {
	cursor, err := repository.Buildings.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var result []entities.Building
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
