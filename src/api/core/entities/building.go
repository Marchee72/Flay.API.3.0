package entities

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Building struct {
	ID      primitive.ObjectID `json:"ID" bson:"ID"`
	Name    string             `json:"name" bson:"name"`
	Address Address            `json:"address" bson:"address"`
}

func (building Building) ToLw() lw.BuildingLw {
	return lw.NewBuildingLw(building.ID, building.Name)
}
