package cmd

import (
	"context"
	"fmt"
	"log"

	"example.com/nestor-expense-tracker/expenses"
	"example.com/nestor-expense-tracker/misc"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	filterMonth    string
	filterCategory string
	filterStore    string
)

var summarizeExpenses = &cobra.Command{
	Use:   "summarize",
	Short: "Summarize the info relating to expenses based on flags provided",
	Run: func(cmd *cobra.Command, args []string) {
		// declare prompter
		prompter := misc.Prompter{}
		prompter.Init()
		// declare filter
		filter := bson.M{}
		// get expenses collection
		expensesCollection := expenses.ExpensesCollection()
		// format date if it is defined and provided
		if filterMonth != "" {
			for {
				startDate, endDate, err := misc.GetMonthRange(filterMonth)
				if err != nil {
					fmt.Println(err.Error())
					filterMonth = prompter.PromptUserFreeForm("What is the month you were wanting to get summary info on? (MM-YY format):")
					continue
				}

				filter["date"] = bson.M{
					"$gte": startDate,
					"$lte": endDate,
				}
				break
			}
		}

		// format category if defined and provided
		if filterCategory != "" {
			// see if category is in list of them
			for !expenses.IsExpenseCategory(filterCategory) {
				userInput := prompter.PromptUserOptions("Select the kind of category you want a summary for:", expenses.GetExpenseCategoryNames())
				categorySelected, err := expenses.GetExpenseCategoryByIndex(userInput)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				filterCategory = categorySelected
			}

			filter["category"] = filterCategory
		}

		// add store to search query if provided
		if store != "" {
			filter["store"] = store
		}

		// make the query
		cursor, err := expensesCollection.Find(context.TODO(), filter)
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer cursor.Close(context.TODO())

		totalSpent, totalTxs := 0.00, 0
		for cursor.Next(context.TODO()) {
			var result expenses.Expense
			if err := cursor.Decode(&result); err != nil {
				log.Printf("Decode error:%s\n", err)
				continue
			}
			fmt.Println(result)
			totalSpent += result.Amount
			totalTxs += 1
		}
		totalSpentStr, err := misc.FormatCurrency(totalSpent)
		if err != nil {
			log.Fatalf("Error formatting total spent: %s\n\n", err.Error())
		}

		fmt.Printf("IN TOTAL: %s over %d transactions", totalSpentStr, totalTxs)

		if err := cursor.Err(); err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func init() {
	summarizeExpenses.Flags().StringVarP(&filterMonth, "month", "m", "", "Specify the month to get a summary for (MM-YY)")
	summarizeExpenses.Flags().StringVarP(&filterCategory, "category", "c", "", "Specify the category to get a summary for")
	summarizeExpenses.Flags().StringVarP(&filterStore, "store", "s", "", "Specify the store in which your summary will apply to")

	RootCmd.AddCommand(summarizeExpenses)
}
