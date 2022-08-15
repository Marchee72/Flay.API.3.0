package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	get_building_booking_contracts "flay-api-v3.0/src/api/core/contracts/get_building_bookings"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/usecases/get_building_bookings"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetBuildingBookings struct {
	GetUserBookingsUseCase get_building_bookings.UseCase
}

func (handler *GetBuildingBookings) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter, constants.UserAdmin})
}

func (handler *GetBuildingBookings) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()
	claims, err := authentication.GetUserCredentials(c)
	log.Printf(claims.Username)
	if err != nil {
		log.Printf("Error authenticating user: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	request := get_building_booking_contracts.Request{}
	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	if err := c.BindQuery(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	response, err := handler.GetUserBookingsUseCase.Execute(ctx, request)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, response)
	return nil
}
