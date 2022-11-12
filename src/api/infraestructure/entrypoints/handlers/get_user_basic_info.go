package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/usecases/get_user_basic_info"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetUserBasicInfo struct {
	GetUserBasicInfoUseCase get_user_basic_info.UseCase
}

func (handler *GetUserBasicInfo) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, constants.AnyUserType())
}

func (handler *GetUserBasicInfo) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()

	claims, err := authentication.GetUserCredentials(c)
	if err != nil {
		log.Printf("Error authenticating user: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	response, err := handler.GetUserBasicInfoUseCase.Execute(ctx, claims.ID)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, response)
	return nil
}
