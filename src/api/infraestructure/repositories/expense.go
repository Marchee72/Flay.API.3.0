package store

import (
	"context"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpenseRepository struct {
	Expenses *mongo.Collection
}

func (repository *ExpenseRepository) SaveExpense(ctx context.Context, expense entities.Expense) error {
	_, err := repository.Expenses.InsertOne(ctx, expense)
	return err
}

func (repository *ExpenseRepository) GetExpensesByUnit(ctx context.Context, unit entities.Unit) ([]entities.Expense, error) {
	cursor, err := repository.Expenses.Find(ctx, bson.M{"unit.floor": unit.Floor, "unit.apartment": unit.Apartment})
	if err != nil {
		return nil, err
	}
	var result []entities.Expense
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
