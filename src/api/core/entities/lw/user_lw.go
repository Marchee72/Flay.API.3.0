package lw

import (
	"flay-api-v3.0/src/api/core/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLw struct {
	ID   primitive.ObjectID `bson:"id"`
	Name constants.UserType `bson:"name"`
}
