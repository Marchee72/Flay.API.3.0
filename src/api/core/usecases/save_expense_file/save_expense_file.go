package save_expense_file

import (
	"context"
	"io"

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
	file, err := request.File.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	unit, err := entities.NewUnit(request.Unit)
	if err != nil {
		return err
	}

	// Read the file contents into a byte slice
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	expense := entities.NewExpense(request.Filename, lw.NewBuildingLw(request.Building.ID, request.Building.Name), *unit, request.Month, int(request.Year), data)
	return usecase.ExpenseRepository.SaveExpense(ctx, expense)
}
