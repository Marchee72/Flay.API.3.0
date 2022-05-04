package entities

import (
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonSpace struct {
	ID       primitive.ObjectID `bson:"ID"`
	Name     string             `bson:"name"`
	Building lw.BuildingLw      `bson:"buildng"`
	Type     string             `bson:"type"`
}
