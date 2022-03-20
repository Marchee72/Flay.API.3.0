package providers

import "go.mongodb.org/mongo-driver/bson/primitive"

type Authentication interface {
	CreateToken(userId primitive.ObjectID) (string, error)
}
