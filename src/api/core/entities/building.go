package entities

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Building struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Address      Address            `bson:"address"`
	CommonSpaces []CommonSpace      `bson:"common_spaces"`
	Admin        lw.UserLw          `bson:"admin"`
}

func (building Building) ToLw() lw.BuildingLw {
	return lw.NewBuildingLw(building.ID, building.Name)
}
