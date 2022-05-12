package lw

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommonSpaceLw struct {
	ID   primitive.ObjectID `bson:"id"`
	Name string             `bson:"name"`
}
