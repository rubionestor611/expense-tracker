package expenses

import (
	"errors"
)

type ExpenseCategory int

const (
	Mortgage ExpenseCategory = iota
	Food
	Utilities
	Entertainment
	Debt
	Misc
	Home_Goods
)

func (e ExpenseCategory) String() string {
	return [...]string{"Mortgage", "Food", "Utilities", "Entertainment", "Debt", "Misc", "Home_Goods"}[e]
}

func GetExpenseCategories() []ExpenseCategory {
	return []ExpenseCategory{Mortgage, Food, Utilities, Entertainment, Debt, Misc, Home_Goods}
}

func GetExpenseCategoryStrings() []string {
	return []string{"Mortgage", "Food", "Utilities", "Entertainment", "Debt", "Misc", "Home_Goods"}
}

func GetExpenseCategoryByIndex(index int) (string, error) {
	strings := GetExpenseCategoryStrings()
	if index < 0 || index >= len(strings) {
		return "", errors.New("index out of bounds on Expense Category enum")
	}

	return strings[index], nil
}
