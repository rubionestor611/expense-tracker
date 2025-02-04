package cmd

import (
	"fmt"
	"log"
	"strings"

	"example.com/nestor-expense-tracker/expenses"
	"github.com/spf13/cobra"
)

var expenseCmd = &cobra.Command{
	Use:   "expense",
	Short: "Expense is the opening command allowing you to add, remove, list, view items logged so far",
	Run: func(cmd *cobra.Command, args []string) {
		var command string

		if len(args) == 0 {
			fmt.Println("You must select an action like add, remove, list, view")
		} else {
			command = strings.ToLower(args[0])

			if command == "add" {
				addErr := expenses.AddExpense()
				if addErr != nil {
					log.Fatalf("Big issues in the world of adding expenses")
				}
			} else if command == "remove" {
				log.Println("Not supported rn gang")
			} else {
				log.Println("BOYYY")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(expenseCmd)
}
