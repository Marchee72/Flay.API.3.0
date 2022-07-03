package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	book_common_space "flay-api-v3.0/src/api/core/contracts/book_common_space"
	"flay-api-v3.0/src/api/core/contracts/common"
	"flay-api-v3.0/src/api/core/entities/lw"
	"flay-api-v3.0/src/api/core/errors"
	bookCommonSpaceUseCase "flay-api-v3.0/src/api/core/usecases/book_common_space"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/authentication"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type BookCommonSpace struct {
	BookCommonSpaceUseCase bookCommonSpaceUseCase.UseCase
}

func (handler *BookCommonSpace) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter})
}

func (handler *BookCommonSpace) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()

	request := book_common_space.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	claims, err := authentication.GetUserCredentials(c)
	if err != nil {
		log.Printf("Error authenticating user: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	request.User = common.NewUserLw(lw.UserLw{ID: claims.ID, Username: claims.Username})
	response, err := handler.BookCommonSpaceUseCase.Execute(ctx, request)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}

	c.JSON(http.StatusOK, response)
	return nil
}
