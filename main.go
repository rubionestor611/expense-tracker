/*
Copyright © 2025 nestordrubio9@gmail.com
*/
package main

import (
	"log"
	"os"

	"example.com/nestor-expense-tracker/cmd"
	"example.com/nestor-expense-tracker/expenses"
	"github.com/joho/godotenv"
)

func main() {
	// load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to mongo
	expenses.InitMongo(os.Getenv("MONGO_CONNECTION_STR"))

	cmd.Execute()
}
