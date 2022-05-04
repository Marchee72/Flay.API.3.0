package authentication

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/entities"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"
)

const BEARER_SCHEMA = "Bearer:"

func CreateToken(user entities.User) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := CustomClaims{
		ID:         user.ID,
		UserType:   user.UserType,
		Authorized: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	token = BEARER_SCHEMA + token
	msg := fmt.Sprintf("creating credentials for user_id: &s, user_type: %s", user.ID, user.UserType)
	fmt.Println(msg)
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
		return nil, errors.New("token does not implements bearer schema.")
	}
	token, err := VerifySigning(tokenString[1])
	if !token.Valid {
		return nil, errors.New("invalid authentication token.")
	}
	if err != nil {
		return nil, err
	}
	return token, nil
}

func Claims(c *gin.Context) (*CustomClaims, error) {
	token, err := ExtractToken(c)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(CustomClaims); ok {
		return &claims, nil
	}
	return nil, fmt.Errorf("an error ocurred trying to claim permissions.")
}
