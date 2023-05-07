package get_unit_expenses

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/get_expenses"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, unit string) (*get_expenses.Response, error)
}

type Implementation struct {
	ExpensesRepository providers.ExpenseRepository
}

func (usecase Implementation) Execute(ctx context.Context, unit string) (*get_expenses.Response, error) {
	u, err := entities.NewUnit(unit)
	if err != nil {
		return nil, err
	}
	expenses, err := usecase.ExpensesRepository.GetExpensesByUnit(ctx, *u)
	if err != nil {
		return nil, err
	}
	response := get_expenses.Response{}
	for _, e := range expenses {
		response.AddExpense(e)
	}
	return &response, nil
}
