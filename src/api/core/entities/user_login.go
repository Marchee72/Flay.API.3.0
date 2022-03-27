package entities

import (
	"flay-api-v3.0/src/api/core/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLogin struct {
	ID       primitive.ObjectID `json:"ID" bson:"_id"`
	UserType constants.UserType `json:"type" bson:"type"`
}
