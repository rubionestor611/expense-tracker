package expenses

import (
	"context"
	"fmt"

	"example.com/nestor-expense-tracker/misc"
)

func addToMongo(expense Expense) error {
	mongoClient := GetMongoClient()

	expenseCollection := mongoClient.Database("expenses").Collection("expenses")

	_, err := expenseCollection.InsertOne(context.TODO(), expense)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully inserted expense %v", expense)
	return nil
}

func AddExpense() (err error) {
	options := GetExpenseCategoryStrings()
	prompter := misc.Prompter{}
	prompter.Init()

	// get date of transaction
	date, err := misc.GetTimeInTimezone("EST")
	if err != nil {
		fmt.Println("Oops! Looks like we can't get today's date. Nestor needs to look into this...")
		return err
	}
	formattedDate := misc.ISOFormat(*date)

	// get category of transaction
	categoryIndex := prompter.PromptUserOptions("Select the type of purchase:", options)
	categorySelection, err := GetExpenseCategoryByIndex(categoryIndex)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// get the store the purchase was at
	var purchaseStore string
	if categoryIndex == int(Mortgage) {
		purchaseStore = "Mortgage Holder"
	} else {
		purchaseStore = prompter.PromptUserFreeForm("What store was the purchase in?")
	}

	// get the amount of the transaction
	var purchaseAmount float64
	var formattedPurchaseAmt string
	for {
		purchaseAmount = prompter.PromptUserFloat("How much was the transaction?", true)

		if purchaseAmount > 0 {
			formattedPurchaseAmt, err = misc.FormatCurrency(purchaseAmount)

			if err != nil {
				return err
			}
			break
		}

		fmt.Printf("You must provide a positive amount for the purchase you made. Please try again\n\n")
	}

	fmt.Println(formattedDate, categorySelection, purchaseStore, formattedPurchaseAmt)

	err = addToMongo(Expense{Date: formattedDate, Category: categorySelection, Store: purchaseStore, Amount: purchaseAmount})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
