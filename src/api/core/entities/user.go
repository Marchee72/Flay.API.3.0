package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"ID" bson:"ID"`
	UserName string             `json:"username" bson:"username"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
}
