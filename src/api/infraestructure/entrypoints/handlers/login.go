package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
}

func (handler *Login) Handle(c *gin.Context) {

	// err := handler.ProcessFileUseCase.Execute(ctx)
	// if err != nil {
	// 	return errors.GetCommonsAPIError(err)
	// }

	c.JSON(http.StatusOK, nil)
}
