package lw

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuildingLw struct {
	ID   primitive.ObjectID `json:"ID" bson:"ID"`
	Name string             `json:"name" bson:"name"`
}

func NewBuildingLw(id primitive.ObjectID, name string) BuildingLw {
	return BuildingLw{
		ID:   id,
		Name: name,
	}
}
