package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/contracts/login"
	loginUseCase "flay-api-v3.0/src/api/core/usecases"

	"github.com/gin-gonic/gin"
)

type Login struct {
	LoginUseCase loginUseCase.UseCase
}

func (handler *Login) Handle(c *gin.Context) {
	ctx := context.TODO()

	request := login.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	result, err := handler.LoginUseCase.Execute(ctx, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}
