package entities

import (
	"flay-api-v3.0/src/api/core/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Credentials struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	UserType constants.UserType `bson:"type"`
}
