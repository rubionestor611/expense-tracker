package expenses

type Expense struct {
	Date     string  `bson:"date"`
	Category string  `bson:"category"`
	Store    string  `bson:"store"`
	Amount   float64 `bson:"amount"`
}
