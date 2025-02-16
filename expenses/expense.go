package expenses

import "time"

type Expense struct {
	Date     time.Time `bson:"date"`
	Category string    `bson:"category"`
	Store    string    `bson:"store"`
	Amount   float64   `bson:"amount"`
}
