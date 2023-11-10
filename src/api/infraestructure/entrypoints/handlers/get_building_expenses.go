package handlers

import (
	"context"
	"log"
	"net/http"

	"flay-api-v3.0/src/api/core/constants"
	get_building_expenses_contract "flay-api-v3.0/src/api/core/contracts/get_building_expenses"
	"flay-api-v3.0/src/api/core/errors"
	"flay-api-v3.0/src/api/core/usecases/get_building_expenses"
	"flay-api-v3.0/src/api/infraestructure/api_errors"
	"flay-api-v3.0/src/api/infraestructure/entrypoints/wrappers"
	"github.com/gin-gonic/gin"
)

type GetBuildingExpenses struct {
	GetBuildingExpensesUseCase get_building_expenses.UseCase
}

func (handler *GetBuildingExpenses) Handle(c *gin.Context) {
	wrappers.AuthWrapper(handler.handle, c, []constants.UserType{constants.UserAdmin})
}

func (handler *GetBuildingExpenses) handle(c *gin.Context) *api_errors.APIError {
	ctx := context.Background()
	request := get_building_expenses_contract.Request{}
	if err := c.ShouldBindUri(&request); err != nil {
		log.Printf("Error binding uri: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	if err := c.ShouldBindQuery(&request); err != nil {
		log.Printf("Error binding query: %s", err.Error())
		return errors.GetCommonsApiError(err)
	}
	response, err := handler.GetBuildingExpensesUseCase.Execute(ctx, request)
	if err != nil {
		return errors.GetCommonsApiError(err)
	}
	c.JSON(http.StatusOK, response)
	return nil

}
