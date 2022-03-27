package authentication

import (
	"flay-api-v3.0/src/api/core/constants"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomClaims struct {
	ID         primitive.ObjectID `json:"id"`
	UserType   constants.UserType `json:"type"`
	Authorized bool               `json:"authorized"`
	jwt.StandardClaims
}
