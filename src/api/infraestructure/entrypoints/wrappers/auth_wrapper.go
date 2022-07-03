package wrappers

import (
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(ctx *gin.Context) *api_errors.APIError

// ErrorWrapper if handlerFunc return a error,then response will be composed from error's information.
func AuthWrapper(handlerFunc HandlerFunc, c *gin.Context, allowedUsers []constants.UserType) {
	token, err := authentication.ExtractToken(c)
	if err != nil {
		apiErr := errors.GetCommonsApiError(err)
		c.JSON(apiErr.Status, apiErr)
		return
	}
	if !authentication.IsAllowed(token, allowedUsers) {
		c.JSON(http.StatusUnauthorized, "you are not allowed to perform this action.")
		return
	}
	apiErr := handlerFunc(c)
	if err != nil {
		c.JSON(apiErr.Status, apiErr)
		return
	}
}
