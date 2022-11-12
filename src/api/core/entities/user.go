package entities

import (
	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities/lw"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	UserName  string             `json:"username" bson:"username"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	UserType  constants.UserType `json:"type" bson:"type"`
	Building  *lw.BuildingLw     `json:"building" bson:"building"`
	Apartment *lw.ApartmentLw    `json:"apartment" bson:"apartment"`
}

func (u User) ToLw() lw.UserLw {
	return lw.UserLw{ID: u.ID, Username: u.UserName}
}
