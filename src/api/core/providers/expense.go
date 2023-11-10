package providers

import (
	"context"
	"time"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExpenseRepository interface {
	SaveExpense(ctx context.Context, expense entities.Expense) error
	GetExpensesByUnit(ctx context.Context, unit entities.Unit) ([]entities.Expense, error)
	GetExpensesByBuilding(ctx context.Context, buildingID primitive.ObjectID, month time.Month, year int) ([]entities.Expense, error)
}
