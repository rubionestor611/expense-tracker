package categories

import (
	"fmt"
	"log"
	"strings"

	"example.com/expense-tracker/expenses"
	"example.com/expense-tracker/misc"
	"github.com/spf13/cobra"
)

var categoryName string

var removeCategoryCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove an expense category",
	Run: func(cmd *cobra.Command, args []string) {
		prompter := misc.Prompter{}

		if strings.EqualFold(categoryName, "") {
			categoryName = prompter.PromptUserFreeForm("Provide the name of the spending category you wish to remove")
		}

		for {
			success, err := expenses.DeleteExpenseCategory(categoryName)
			if err != nil {
				log.Fatal("Error deleting a category", err)
			}
			if success {
				fmt.Printf("Successfully deleted expense category %s\n\n", categoryName)
				break
			} else {
				categoryName = prompter.PromptUserFreeForm(fmt.Sprintf("%s was not a category to delete. Please provide an existing one", categoryName))
			}
		}
	},
}

func init() {
	removeCategoryCommand.Flags().StringVarP(&categoryName, "name", "n", "", "Define the name of category to remove")
}
