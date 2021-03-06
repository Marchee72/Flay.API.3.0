package authentication

import (
	"fmt"
	"log"
	"os"
	"strings"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/errors"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

const BEARER_SCHEMA = "Bearer:"

func CreateToken(user entities.Credentials) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := CustomClaims{
		ID:         user.ID,
		UserType:   user.UserType,
		Username:   user.Username,
		Authorized: true,
		// StandardClaims: jwt.StandardClaims{
		// 	ExpiresAt: time.Now().Add(time.Minute * 150).Unix(),
		// },
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	token = BEARER_SCHEMA + token
	log.Printf("creating credentials for user_id: %s, user_name: %s user_type: %s", user.ID, user.Username, user.UserType)
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
	if len(allowedUsers) == 0 {
		return true
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userType := claims["type"]
		for i := range allowedUsers {
			if allowedUsers[i].Equals(userType) {
				return true
			}
		}
		fmt.Printf("user not allowed for this action. User type: %s.", userType)
	} else if !ok {
		fmt.Printf("an error ocurred trying to claim permissions.")
	}
	return false
}

func ExtractToken(c *gin.Context) (*jwt.Token, error) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Split(authHeader, BEARER_SCHEMA)
	if len(tokenString) != 2 {
		return nil, errors.NewUnauthorizedError("token does not implements bearer schema")
	}
	token, err := VerifySigning(tokenString[1])
	if err != nil || !token.Valid {
		return nil, errors.NewUnauthorizedError("invalid authentication token")
	}
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetUserCredentials(c *gin.Context) (*entities.Credentials, error) {
	token, err := ExtractToken(c)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)
	s := GetCredentialsFromClaims(claims)
	return &s, nil
}
