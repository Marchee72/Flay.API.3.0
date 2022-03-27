package handlers

import (
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	bookCommonAreaUseCase "flay-api-v3.0/src/api/core/usecases/book_common_area"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type BookCommonArea struct {
	BookCommonAreaUseCase bookCommonAreaUseCase.UseCase
}

func (handler *BookCommonArea) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserAdmin})
}

func (handler *BookCommonArea) handle(c *gin.Context) error {
	c.JSON(http.StatusOK, "SUUUUUUUU!")
	return nil
}
