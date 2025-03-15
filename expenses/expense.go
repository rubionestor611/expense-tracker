package expenses

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	ID       primitive.ObjectID `bson:"_id"`
	Date     time.Time          `bson:"date"`
	Category string             `bson:"category"`
	Store    string             `bson:"store"`
	Amount   float64            `bson:"amount"`
}

func (expense Expense) String() string {
	return fmt.Sprintf("{ id: %s, date: %s, category: %s, store: %s, amount: %.2f}", expense.ID.Hex(), expense.Date.Format("01-02-2006"), expense.Category, expense.Store, expense.Amount)
}
