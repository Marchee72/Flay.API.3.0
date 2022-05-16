package lw

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLw struct {
	ID       primitive.ObjectID `bson:"id"`
	Username string             `bson:"name"`
}
