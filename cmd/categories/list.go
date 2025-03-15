package categories

import (
	"fmt"
	"log"

	"example.com/nestor-expense-tracker/expenses"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all expense categories which can be used",
	Run: func(cmd *cobra.Command, args []string) {
		allCategories, err := expenses.GetExpenseCategories()
		if err != nil {
			log.Fatal(err.Error())
		}

		for _, cat := range allCategories {
			fmt.Println(cat)
		}
	},
}

func init() {

}
