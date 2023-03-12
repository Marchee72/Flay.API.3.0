package lw

import (
	"flay-api-v3.0/src/api/core/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLw struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"name"`
	Type     constants.UserType `bson:"type"`
}
