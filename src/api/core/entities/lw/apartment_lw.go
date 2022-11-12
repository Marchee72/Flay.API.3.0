package lw

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApartmentLw struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func NewApartmentLw(id primitive.ObjectID, floor uint8, flat string) ApartmentLw {
	name := fmt.Sprintf("%d-%s", floor, flat)
	return ApartmentLw{
		ID:   id,
		Name: name,
	}

}
