package authentication

import (
	"errors"
	"fmt"
	"os"
	"time"

	"flay-api-v3.0/src/api/core/constants"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateToken(userId primitive.ObjectID) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["user_type"] = ""
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifySigning(receivedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func IsAllowed(token *jwt.Token, allowedUsers []constants.UserType) bool {
	if allowedUsers == nil || len(allowedUsers) == 0 {
		return true
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userType := claims["user_type"]
		for i := range allowedUsers {
			if allowedUsers[i] == userType {
				return true
			}
		}
		fmt.Printf("user not allowed for this action. User type: %s", userType)
	} else if !ok {
		fmt.Printf("an error ocurred trying to claim permissions")
	}
	return false
}

func ExtractToken(ctx *gin.Context) (*jwt.Token, error) {
	const BEARER_SCHEMA = "Bearer"
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := VerifySigning(tokenString)
	if !token.Valid {
		return nil, errors.New("invalid authentication token.")
	}
	if err != nil {
		return nil, err
	}
	return token, nil
}
