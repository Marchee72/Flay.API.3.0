package providers

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
)

type ExpenseRepository interface {
	SaveExpense(ctx context.Context, expense entities.Expense) error
	GetExpensesByUnit(ctx context.Context, unit entities.Unit) ([]entities.Expense, error)
}
