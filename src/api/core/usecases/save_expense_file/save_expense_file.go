package save_expense_file

import (
	"context"

	"flay-api-v3.0/src/api/core/contracts/save_expense_file"
	"flay-api-v3.0/src/api/core/entities"
	"flay-api-v3.0/src/api/core/entities/lw"
	"flay-api-v3.0/src/api/core/providers"
)

type UseCase interface {
	Execute(ctx context.Context, request save_expense_file.Request) error
}

type Implementation struct {
	ExpenseRepository providers.ExpenseRepository
}

func (usecase Implementation) Execute(ctx context.Context, request save_expense_file.Request) error {
	renter := lw.UserLw{
		ID:       request.Renter.ID,
		Username: request.Renter.Name,
	}
	expense := entities.NewExpense(request.Filename, renter, request.Month, int(request.Year), request.File)
	return usecase.ExpenseRepository.SaveExpense(ctx, expense)
}
