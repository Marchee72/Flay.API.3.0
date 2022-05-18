package entities

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Building struct {
	ID           primitive.ObjectID `bson:"ID"`
	Name         string             `bson:"name"`
	Address      Address            `bson:"address"`
	CommonSpaces []CommonSpace      `bson:"common_spaces"`
}

func (building Building) ToLw() lw.BuildingLw {
	return lw.NewBuildingLw(building.ID, building.Name)
}
