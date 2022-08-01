package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/usecases/get_user_building"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetUserBuilding struct {
	GetUserBuildingUseCase get_user_building.UseCase
}

func (handler *GetUserBuilding) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter, constants.UserAdmin})
}

func (handler *GetUserBuilding) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()
	claims, err := authentication.GetUserCredentials(c)
	log.Printf(claims.Username)
	if err != nil {
		log.Printf("Error authenticating user: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	response, err := handler.GetUserBuildingUseCase.Execute(ctx, claims.ID)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, response)
	return nil
}
