package expenses

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
