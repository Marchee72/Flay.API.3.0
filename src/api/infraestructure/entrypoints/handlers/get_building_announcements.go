package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	get_building_announcements_contract "flay-api-v3.0/src/api/core/contracts/get_building_announcements"
	"flay-api-v3.0/src/api/core/errors"

	"flay-api-v3.0/src/api/core/usecases/get_building_announcements"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetBuildingAnnouncements struct {
	GetBuildingAnnouncementsUseCase get_building_announcements.UseCase
}

func (handler *GetBuildingAnnouncements) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter, constants.UserAdmin})
}

func (handler *GetBuildingAnnouncements) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()
	request := get_building_announcements_contract.Request{}
	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("Error binding uri: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	if err := c.BindQuery(&request); err != nil {
		log.Printf("Error binding query: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	response, err := handler.GetBuildingAnnouncementsUseCase.Execute(ctx, request)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, response)
	return nil
}
