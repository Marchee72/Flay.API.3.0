package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	get_announcement_contract "flay-api-v3.0/src/api/core/contracts/get_announcement"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/usecases/get_announcement"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetAnnouncement struct {
	GetAnnouncementUseCase get_announcement.UseCase
}

func (handler *GetAnnouncement) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter, constants.UserAdmin})
}

func (handler *GetAnnouncement) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()
	request := get_announcement_contract.Request{}
	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("Error binding uri: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	response, err := handler.GetAnnouncementUseCase.Execute(ctx, request)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, response)
	return nil
}
