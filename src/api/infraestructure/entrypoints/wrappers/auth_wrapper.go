package wrappers

import (
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(ctx *gin.Context) error

// ErrorWrapper if handlerFunc return a error,then response will be composed from error's information.
func AuthWrapper(handlerFunc HandlerFunc, ctx *gin.Context, allowedUsers []constants.UserType) {
	token, err := authentication.ExtractToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if !authentication.IsAllowed(token, allowedUsers) {
		ctx.JSON(http.StatusUnauthorized, "you are not allowed to perform this action.")
	}
	err = handlerFunc(ctx)
}
