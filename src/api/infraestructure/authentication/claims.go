package authentication

import (
	"flay-api-v3.0/src/api/core/constants"
	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomClaims struct {
	ID         primitive.ObjectID `json:"id"`
	UserType   constants.UserType `json:"type"`
	Name       string             `json:"name"`
	Authorized bool               `json:"authorized"`
	jwt.StandardClaims
}
