package store

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApartmentRepository struct {
	Apartments *mongo.Collection
}

func (repository *ApartmentRepository) GetApartment(ctx context.Context, apartmentID primitive.ObjectID) (*entities.Apartment, error) {
	var result entities.Apartment
	if err := repository.Apartments.FindOne(ctx, bson.M{"_id": apartmentID}).Decode(&result); err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	return &result, nil
}
