package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"example.com/expense-tracker/expenses"
	"example.com/expense-tracker/misc"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		allCategories, err := expenses.GetExpenseCategories()
		if err != nil {
			log.Fatalf(err.Error())
		}
		categoryNames := expenses.ExtractCategoryNames(allCategories)

		submittingExpense := expenses.Expense{
			ID:       primitive.NewObjectID(),
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

			categoryIndex := prompter.PromptUserOptions("Select the type of purchase:", categoryNames)
			categorySelection, err := expenses.GetExpenseCategoryByIndex(categoryIndex)
			if err != nil {
				continue
			} else {
				submittingExpense.Category = categorySelection
				break
			}
		}

		// get the category. if the category has a defaultStore use that
		for _, cat := range allCategories {
			if strings.EqualFold(cat.Name, submittingExpense.Category) {
				if len(cat.DefaultStore) > 0 {
					submittingExpense.Store = cat.DefaultStore
				}
				break
			}
		}

		for strings.EqualFold(submittingExpense.Store, "") {
			submittingExpense.Store = strings.ToLower(prompter.PromptUserFreeForm("What store was the purchase in?"))
		}

		expenseCollection := expenses.ExpensesCollection()

		_, err = expenseCollection.InsertOne(context.TODO(), submittingExpense)
		if err != nil {
			log.Fatalf("Error adding expense: %v", err)
		}

		fmt.Printf("Successfully submitted expense\n%s\n\n", submittingExpense)
	},
}

func init() {
	addExpenseCmd.Flags().BoolVarP(&today, "today", "t", false, "Mark the expense as having happened today")
	addExpenseCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Specify the amount of the transaction")
	addExpenseCmd.Flags().StringVarP(&category, "category", "c", "", "Specify the category that the transaction falls under. If no match found, you will be prompted again")
	addExpenseCmd.Flags().StringVarP(&store, "store", "s", "", "Specify the store in which the expense was made")

	RootCmd.AddCommand(addExpenseCmd)
}
