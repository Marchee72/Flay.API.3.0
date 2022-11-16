package common

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apartment struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}

func NewApartment(a lw.ApartmentLw) Apartment {
	return Apartment{
		ID:   a.ID,
		Name: a.Name,
	}
}
