package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApartmentRepository interface {
	GetApartment(ctx context.Context, apartmentID primitive.ObjectID) (*entities.Apartment, error)
}
