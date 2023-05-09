package store

import (
	"context"
	"fmt"

	"flay-api-v3.0/src/api/core/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExpenseRepository struct {
	Expenses *mongo.Collection
	Files    *FileRepository
}

func (repository *ExpenseRepository) SaveExpense(ctx context.Context, expense entities.Expense) error {
	result, err := repository.Expenses.InsertOne(ctx, expense)
	if err != nil {
		return err
	}
	contentID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("error parsing inserted id: %s", contentID)
	}
	return repository.Files.SaveFile(ctx, expense.File, contentID)
}

func (repository *ExpenseRepository) GetExpensesByUnit(ctx context.Context, unit entities.Unit) ([]entities.Expense, error) {
	cursor, err := repository.Expenses.Find(ctx, bson.M{"unit.floor": unit.Floor, "unit.apartment": unit.Apartment})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var result []entities.Expense
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
