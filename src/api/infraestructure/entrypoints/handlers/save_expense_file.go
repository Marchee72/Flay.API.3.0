package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	contract "flay-api-v3.0/src/api/core/contracts/save_expense_file"
	"flay-api-v3.0/src/api/core/errors"

	"flay-api-v3.0/src/api/core/usecases/save_expense_file"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type SaveExpenseFile struct {
	SaveExpenseFileUseCase save_expense_file.UseCase
}

func (handler *SaveExpenseFile) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter, constants.UserAdmin})
}

func (handler *SaveExpenseFile) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()

	request := contract.Request{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	if err := handler.SaveExpenseFileUseCase.Execute(ctx, request); err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, nil)
	return nil
}
