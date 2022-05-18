package wrappers

import (
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(ctx *gin.Context)

// ErrorWrapper if handlerFunc return a error,then response will be composed from error's information.
func AuthWrapper(handlerFunc HandlerFunc, c *gin.Context, allowedUsers []constants.UserType) {
	token, err := authentication.ExtractToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if !authentication.IsAllowed(token, allowedUsers) {
		c.JSON(http.StatusUnauthorized, "you are not allowed to perform this action.")
		return
	}
	handlerFunc(c)
}
