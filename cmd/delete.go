package cmd

import (
	"context"
	"fmt"
	"log"

	"example.com/nestor-expense-tracker/expenses"
	"example.com/nestor-expense-tracker/misc"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	id string
)

var deleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete a transaction",
	Run: func(cmd *cobra.Command, args []string) {
		if len(id) == 0 {
			prompter := misc.Prompter{}

			id = prompter.PromptUserFreeForm("Provide the id of the transaction you want to delete:")
		}
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Fatalf(err.Error())
		}

		filter := bson.M{
			"_id": objectId,
		}

		res, err := expenses.ExpensesCollection().DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatal(err.Error())
		}

		if res.DeletedCount == 1 {
			fmt.Printf("Successfully deleted transaction %s!\n\n", id)
		} else {
			fmt.Printf("Uh oh! Seems like %s wasn't the id of an existing expense. Try running `expense-tracker summarize` to see the ids of the expenses", id)
		}
	},
}

func init() {
	deleteCommand.Flags().StringVarP(&id, "id", "i", "", "ID of transaction to delete")
	RootCmd.AddCommand(deleteCommand)
}
