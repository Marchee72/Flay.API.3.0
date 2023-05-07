package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	"flay-api-v3.0/src/api/core/contracts/get_expenses"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/usecases/get_unit_expenses"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetUnitExpenses struct {
	GetUnitExpensesUseCase get_unit_expenses.UseCase
}

func (handler *GetUnitExpenses) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserRenter})
}

func (handler *GetUnitExpenses) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()
	request := get_expenses.Request{}
	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("Error binding request: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	response, err := handler.GetUnitExpensesUseCase.Execute(ctx, request.Unit)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, response)
	return nil
}
