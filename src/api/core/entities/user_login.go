package entities

import "time"

type UserLogin struct {
	UserName  string    `json:"username" bson:"username"`
	Password  string    `json:"-" bson:"-"`
	LastLogin time.Time `json:"last_login" bson:"last_login"`
}
