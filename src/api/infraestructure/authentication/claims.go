package authentication

import (
	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities"
	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomClaims struct {
	ID         primitive.ObjectID `json:"user_id,omitempty"`
	UserType   constants.UserType `json:"type,omitempty"`
	Username   string             `json:"username,omitempty"`
	Authorized bool               `json:"authorized,omitempty"`
	jwt.StandardClaims
}

func GetCredentialsFromClaims(claims jwt.MapClaims) entities.Credentials {
	objID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	userType, err := constants.ToUserType(claims["type"].(string))
	if err != nil {
		panic(err)
	}
	return entities.Credentials{
		ID:       objID,
		UserType: *userType,
		Username: claims["username"].(string),
	}
}
