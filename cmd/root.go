/*
Copyright © 2025 nestordrubio9@gmail.com
*/
package cmd

import (
	"os"

	"example.com/expense-tracker/cmd/categories"
	"example.com/expense-tracker/expenses"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "My personal expense tracker",
	Long:  `A Golang-based Cobra command line interface (CLI) which I built to manage expenses without a fancy UI and so I can feel cool and hackery while I do it.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// load env vars
	godotenv.Load()
	// connect to mongo
	expenses.InitMongo(os.Getenv("EXPENSES_MONGO_CONNECTION_STR"))
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.expense-tracker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.AddCommand(categories.CategoriesCmd)
}
