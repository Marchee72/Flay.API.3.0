package handlers

import (
	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/usecases/get_user_building"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
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
	return api_errors.NewResourceNotFound()
}
