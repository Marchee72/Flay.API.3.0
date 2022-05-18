package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/usecases/get_user_bookings"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetUserBookings struct {
	GetUserBookingsUseCase get_user_bookings.UseCase
}

func (handler *GetUserBookings) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter, constants.UserAdmin})
}

func (handler *GetUserBookings) handle(c *gin.Context) {
	ctx := context.Background()
	claims, err := authentication.GetUserCredentials(c)
	if err != nil {
		log.Printf("Error authenticating user: %s", err.Error())
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	response, err := handler.GetUserBookingsUseCase.Execute(ctx, claims.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, response)
}
