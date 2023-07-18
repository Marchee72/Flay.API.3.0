package handlers

import (
	"context"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/usecases/get_buildings"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetBuildings struct {
	GetBuildingsUseCase get_buildings.GetBuildingUseCase
}

func (handler *GetBuildings) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserAdmin})
}

func (handler *GetBuildings) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()
	claims, err := authentication.GetUserCredentials(c)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}

	response, err := handler.GetBuildingsUseCase.Execute(ctx, claims.ID)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, response)

	return nil
}
