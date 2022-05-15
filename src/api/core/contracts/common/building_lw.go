package common

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuildingLw struct {
	ID   primitive.ObjectID `bson:"ID"`
	Name string             `bson:"name"`
}

func NewBuildingLw(lw lw.BuildingLw) BuildingLw {
	return BuildingLw{
		ID:   lw.ID,
		Name: lw.Name,
	}
}
