package expenses

import (
	"context"
	"fmt"

	"example.com/nestor-expense-tracker/misc"
)

func AddToMongo(expense Expense) error {
	mongoClient := GetMongoClient()

	expenseCollection := mongoClient.Database("expenses").Collection("expenses")

	_, err := expenseCollection.InsertOne(context.TODO(), expense)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully inserted expense\n{\n\tdate: %s,\n\tcategory: %s,\n\tstore: %s,\n\tamount: %.2f\n}", misc.ISOFormat(expense.Date), expense.Category, expense.Store, expense.Amount)
	return nil
}
