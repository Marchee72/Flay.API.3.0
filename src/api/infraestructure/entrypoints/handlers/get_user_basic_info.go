package handlers

import (
	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/usecases/get_user_basic_info"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetUserBasicInfo struct {
	getUserBasicInfoUseCase get_user_basic_info.UseCase
}

func (handler *GetUserBasicInfo) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, constants.AnyUserType())
}

func (handler *GetUserBasicInfo) handle() *api_errors.APIError {

}
