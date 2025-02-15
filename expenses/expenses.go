package expenses

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"example.com/nestor-expense-tracker/misc"
)

func formatCurrency(val any) (string, error) {
	typeOfVal := reflect.ValueOf(val)

	switch typeOfVal.Kind() {
	case reflect.String:
		floatVal, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			fmt.Printf("There was an error formatting your currency of %s\n", val)
			return "", err
		}

		return fmt.Sprintf("$%.2f", floatVal), nil
	case reflect.Float64, reflect.Float32:
		formatted := fmt.Sprintf("$%.2f", val)
		return formatted, nil
	default:
		errorMsg := fmt.Sprintf("Currently unable to format a value of the type %s", typeOfVal.Kind())
		fmt.Println(errorMsg)
		return "", errors.New(errorMsg)
	}
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
	formattedPurchaseAmt, err := formatCurrency(purchaseAmount)

	if err != nil {
		return err
	}

	fmt.Println(date, categorySelection, purchaseStore, formattedPurchaseAmt)

	return nil
}
