package entities

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apartment struct {
	ID       primitive.ObjectID `bson:"_id"`
	Building lw.BuildingLw      `bson:"building"`
	Owner    User               `bson:"owner"`
	Floor    uint8              `bson:"floor"`
	Flat     string             `bson:"flat"`
}

func (apartment Apartment) ToLw() lw.ApartmentLw {
	return lw.NewApartmentLw(apartment.ID, apartment.Floor, apartment.Flat)
}
