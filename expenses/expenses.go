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

	date, err := misc.GetTimeInTimezone("EST")
	if err != nil {
		fmt.Println("Oops! Looks like we can't get today's date. Nestor needs to look into this...")
		errors.New("Couldn't get time")
	}
	categoryIndex := prompter.PromptUserOptions("Select the type of purchase:", options)
	purchaseStore := prompter.PromptUserFreeForm("What store was the purchase in?")

	categorySelection, err := GetExpenseCategoryByIndex(categoryIndex)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(date, categorySelection, purchaseStore)

	return nil
}
