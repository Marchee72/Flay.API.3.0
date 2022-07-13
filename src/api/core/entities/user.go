package entities

import (
	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"ID" bson:"ID"`
	UserName string             `json:"username" bson:"username"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	UserType constants.UserType `json:"type" bson:"type"`
	Building *lw.BuildingLw     `json:"building" bson:"building"`
}

func (u User) ToLw() lw.UserLw {
	return lw.UserLw{ID: u.ID, Username: u.UserName}
}
