package categories

import (
	"fmt"
	"log"

	"example.com/expense-tracker/expenses"
	"example.com/expense-tracker/misc"
	"github.com/spf13/cobra"
)

var (
	name         string
	description  string
	defaultStore string
)

var addCategoryCommand = &cobra.Command{
	Use:   "add",
	Short: "Add an expense category",
	Run: func(cmd *cobra.Command, args []string) {
		prompter := misc.Prompter{}
		submittingCategory := expenses.ExpenseCategory{
			Name:         name,
			Description:  description,
			DefaultStore: defaultStore,
		}

		noName := len(submittingCategory.Name) == 0
		categoryExists := expenses.IsExpenseCategory(submittingCategory.Name)

		for noName || categoryExists {
			if categoryExists {
				fmt.Printf("%s already exists in our categories. Provide a new one.\n\n", submittingCategory.Name)
			}
			submittingCategory.Name = prompter.PromptUserFreeForm("Please provide a name for the category you wish to introduce:")
			noName = len(submittingCategory.Name) == 0
			categoryExists = expenses.IsExpenseCategory(submittingCategory.Name)
		}

		for len(submittingCategory.Description) == 0 {
			submittingCategory.Description = prompter.PromptUserFreeForm("Please describe what this category will be used for:")
		}

		err := expenses.AddExpenseCategory(submittingCategory)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Printf("Successfully added expense category: %s\n\n", submittingCategory)
	},
}

func init() {
	addCategoryCommand.Flags().StringVarP(&name, "name", "n", "", "Specify the name of the category you wish to add")
	addCategoryCommand.Flags().StringVarP(&description, "description", "d", "", "Specify some description of the categories")
	addCategoryCommand.Flags().StringVarP(&defaultStore, "default store", "s", "", "(Optional) A default store if consistent")
}
