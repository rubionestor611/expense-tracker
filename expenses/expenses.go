package expenses

import (
	"errors"
	"fmt"

	"example.com/nestor-expense-tracker/misc"
)

func AddExpense() (err error) {
	options := GetExpenseCategoryStrings()
	prompter := misc.Prompter{}
	prompter.Init()

	// get date of transaction
	date, err := misc.GetTimeInTimezone("EST")
	if err != nil {
		fmt.Println("Oops! Looks like we can't get today's date. Nestor needs to look into this...")
		errors.New("Couldn't get time")
	}

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
	purchaseAmount := prompter.PromptUserFloat("How much was the transaction?", true)

	fmt.Println(date, categorySelection, purchaseStore, purchaseAmount)

	return nil
}
