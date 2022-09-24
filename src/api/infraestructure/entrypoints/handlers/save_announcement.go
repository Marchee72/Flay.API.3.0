package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/contracts/save_announcement"
	"flay-api-v3.0/src/api/core/errors"
	saveAnnouncementUseCase "flay-api-v3.0/src/api/core/usecases/save_announcement"

	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type SaveAnnouncement struct {
	SaveAnnouncementUseCase saveAnnouncementUseCase.UseCase
}

func (handler *SaveAnnouncement) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter, constants.UserAdmin})
}

func (handler *SaveAnnouncement) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()

	request := save_announcement.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}

	claims, err := authentication.GetUserCredentials(c)
	if err != nil {
		log.Printf("Error authenticating user: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	err = handler.SaveAnnouncementUseCase.Execute(ctx, request, claims.ID)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, nil)
	return nil
}
