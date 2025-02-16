package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"example.com/nestor-expense-tracker/expenses"
	"example.com/nestor-expense-tracker/misc"
	"github.com/spf13/cobra"
)

var (
	today    bool
	amount   float64
	category string
	store    string
)

var addExpenseCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense entry",
	Run: func(cmd *cobra.Command, args []string) {
		prompter := misc.Prompter{}
		allCategories := expenses.GetExpenseCategoryStrings()

		submittingExpense := expenses.Expense{
			Date:     time.Now(),
			Amount:   amount,
			Category: category,
			Store:    store,
		}

		if !today {
			for {
				userInput := prompter.PromptUserFreeForm("Enter the date in MM-DD-YYYY format:")
				parsedDate, err := time.Parse("01-02-2006", userInput)
				if err != nil {
					fmt.Printf("There was an error parsing your date. Make sure to provide it in the right format\n\n")
					continue
				}
				submittingExpense.Date = parsedDate
				break
			}
		}

		if amount <= 0 {
			for submittingExpense.Amount == 0 {
				purchaseAmount := prompter.PromptUserFloat("How much was the transaction?", true)

				if purchaseAmount <= 0 {
					fmt.Println("The value of an expense must be > 0. Try again\n\n")
					continue
				}

				submittingExpense.Amount = purchaseAmount
			}
		}

		// see if in category
		for !expenses.IsExpenseCategory(submittingExpense.Category) {
			if submittingExpense.Category != "" {
				fmt.Printf("You must provide a category for the expense from the available options\n\n")
			}

			categoryIndex := prompter.PromptUserOptions("Select the type of purchase:", allCategories)
			categorySelection, err := expenses.GetExpenseCategoryByIndex(categoryIndex)
			if err != nil {
				continue
			} else {
				submittingExpense.Category = categorySelection
				break
			}
		}

		if strings.EqualFold(submittingExpense.Category, expenses.Mortgage.String()) {
			submittingExpense.Store = "mortgage holder"
		}

		for strings.EqualFold(submittingExpense.Store, "") {
			submittingExpense.Store = strings.ToLower(prompter.PromptUserFreeForm("What store was the purchase in?"))
		}

		addErr := expenses.AddToMongo(submittingExpense)
		if addErr != nil {
			log.Fatalf("Error adding expense: %v", addErr)
		}
	},
}

func init() {
	addExpenseCmd.Flags().BoolVarP(&today, "today", "t", false, "Mark the expense as having happened today")
	addExpenseCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Specify the amount of the transaction")
	addExpenseCmd.Flags().StringVarP(&category, "category", "c", "", "Specify the category that the transaction falls under. If no match found, you will be prompted again")
	addExpenseCmd.Flags().StringVarP(&store, "store", "s", "", "Specify the store in which the expense was made")

	rootCmd.AddCommand(addExpenseCmd)
}
