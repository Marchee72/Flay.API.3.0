package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	book_common_space "flay-api-v3.0/src/api/core/contracts/book_common_area"
	bookCommonAreaUseCase "flay-api-v3.0/src/api/core/usecases/book_common_space"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type BookCommonArea struct {
	BookCommonAreaUseCase bookCommonAreaUseCase.UseCase
}

func (handler *BookCommonArea) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter})
}

func (handler *BookCommonArea) handle(c *gin.Context) {
	ctx := context.Background()

	request := book_common_space.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	claims, err := authentication.Claims(c)
	if err != nil {
		log.Printf("Error authenticating user: %s", err.Error())
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	request.UserID = claims.ID
	response, err := handler.BookCommonAreaUseCase.Execute(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, response)
	return
}
