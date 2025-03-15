package expenses

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type ExpenseCategory struct {
	Name         string `bson:"name"`
	Description  string `bson:"description"`
	DefaultStore string `bson:"defaultStore"`
}

func (e ExpenseCategory) String() string {
	if len(e.DefaultStore) == 0 {
		return fmt.Sprintf("{Name: %s, Description: %s}", e.Name, e.Description)
	}
	return fmt.Sprintf("{Name: %s, Description: %s, DefaultStore: %s}", e.Name, e.Description, e.DefaultStore)
}

func GetExpenseCategories() ([]ExpenseCategory, error) {
	cursor, err := CategoriesCollection().Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

	var results []ExpenseCategory

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func ExtractCategoryNames(categories []ExpenseCategory) []string {
	categoryNames := make([]string, len(categories))
	for index, category := range categories {
		categoryNames[index] = category.Name
	}
	return categoryNames
}

func GetExpenseCategoryNames() []string {
	categories, err := GetExpenseCategories()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return ExtractCategoryNames(categories)
}

func IsExpenseCategory(category string) bool {
	for _, validCategory := range GetExpenseCategoryNames() {
		if strings.EqualFold(validCategory, category) {
			return true
		}
	}
	return false
}

func GetExpenseCategoryByIndex(index int) (string, error) {
	strings := GetExpenseCategoryNames()

	if index < 0 || index >= len(strings) {
		return "", errors.New("index out of bounds on Expense Category enum")
	}

	return strings[index], nil
}

func AddExpenseCategory(cat ExpenseCategory) error {
	_, err := CategoriesCollection().InsertOne(context.TODO(), cat)
	if err != nil {
		return err
	}

	return nil
}
