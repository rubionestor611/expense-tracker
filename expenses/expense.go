package expenses

import (
	"fmt"
	"time"
)

type Expense struct {
	Date     time.Time `bson:"date"`
	Category string    `bson:"category"`
	Store    string    `bson:"store"`
	Amount   float64   `bson:"amount"`
}

func (expense Expense) String() string {
	return fmt.Sprintf("{ date: %s, category: %s, store: %s, amount: %.2f}", expense.Date.Format("01-02-2006"), expense.Category, expense.Store, expense.Amount)
}
