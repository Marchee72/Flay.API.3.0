package get_building_expenses

import (
	"context"
	"time"

	get_building_expenses_contract "flay-api-v3.0/src/api/core/contracts/get_building_expenses"
	"flay-api-v3.0/src/api/core/contracts/get_expenses"

	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request get_building_expenses_contract.Request) (*get_expenses.Response, error)
}

type Implementation struct {
	ExpenseRepository providers.ExpenseRepository
}

func (usecase *Implementation) Execute(ctx context.Context, request get_building_expenses_contract.Request) (*get_expenses.Response, error) {
	if request.Month == nil {
		*request.Month = time.Now().Month()
	}
	if request.Year == nil {
		*request.Year = time.Now().Year()
	}
	expenses, err := usecase.ExpenseRepository.GetExpensesByBuilding(ctx, request.BuildingID(), *request.Month, *request.Year)
	if err != nil {
		return nil, err
	}
	response := get_expenses.Response{}

	for _, e := range expenses {
		response.AddExpense(e)
	}
	return &response, nil

}
