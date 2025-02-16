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
	Home_Goods
	Misc
	Nestor_Personal
	Alyssa_Personal
	Oliver_Personal
)

func (e ExpenseCategory) String() string {
	return [...]string{"Mortgage", "Food", "Utilities", "Entertainment", "Debt", "Home_Goods", "Misc", "Nestor_Personal", "Alyssa_Personal", "Oliver_Personal"}[e]
}

func GetExpenseCategories() []ExpenseCategory {
	return []ExpenseCategory{Mortgage, Food, Utilities, Entertainment, Debt, Home_Goods, Misc}
}

func GetExpenseCategoryStrings() []string {
	return []string{"Mortgage", "Food", "Utilities", "Entertainment", "Debt", "Home_Goods", "Misc", "Nestor_Personal", "Alyssa_Personal", "Oliver_Personal"}
}

func GetExpenseCategoryByIndex(index int) (string, error) {
	strings := GetExpenseCategoryStrings()
	if index < 0 || index >= len(strings) {
		return "", errors.New("index out of bounds on Expense Category enum")
	}

	return strings[index], nil
}

func IsExpenseCategory(category string) bool {
	for _, validCategory := range GetExpenseCategoryStrings() {
		if validCategory == category {
			return true
		}
	}
	return false
}
