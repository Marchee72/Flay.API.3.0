package store

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpenseRepository struct {
	Expenses *mongo.Collection
}

func (repository *ExpenseRepository) SaveExpense(ctx context.Context, expense entities.Expense) error {
	_, err := repository.Expenses.InsertOne(ctx, expense)
	return err
}
