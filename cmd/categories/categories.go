package categories

import (
	"github.com/spf13/cobra"
)

var CategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Info on expense categories",
}

func init() {
	CategoriesCmd.AddCommand(addCategoryCommand)
	CategoriesCmd.AddCommand(listCommand)
	CategoriesCmd.AddCommand(removeCategoryCommand)
}
