package lw

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuildingLw struct {
	ID   primitive.ObjectID `bson:"ID"`
	Name string             `bson:"name"`
}

func NewBuildingLw(id primitive.ObjectID, name string) BuildingLw {
	return BuildingLw{
		ID:   id,
		Name: name,
	}
}
