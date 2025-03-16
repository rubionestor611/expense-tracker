package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
	"unicode/utf8"

	"example.com/expense-tracker/expenses"
	"example.com/expense-tracker/misc"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	filterMonth    string
	filterCategory string
	filterStore    string
	export         bool
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

		var queriedExpenses []expenses.Expense
		err = cursor.All(context.TODO(), &queriedExpenses)
		if err != nil {
			log.Fatalf("Error parsing queried expenses:", err.Error())
			return
		}

		totalSpent, totalTxs := 0.00, 0

		for _, expense := range queriedExpenses {
			fmt.Println(expense)
			totalSpent += expense.Amount
			totalTxs += 1
		}

		totalSpentStr, err := misc.FormatCurrency(totalSpent)
		if err != nil {
			log.Fatalf("Error formatting total spent: %s\n\n", err.Error())
		}

		fmt.Printf("IN TOTAL: %s over %d transactions\n\n", totalSpentStr, totalTxs)

		if err := cursor.Err(); err != nil {
			log.Fatalf(err.Error())
		}

		if export {
			if len(queriedExpenses) == 0 {
				fmt.Println("No expenses to export in this query. Skipping export.")
				return
			}

			tmpDir := os.TempDir()
			tmpFile := filepath.Join(tmpDir, "expense_summary.xlsx")

			defer func() {
				fmt.Println("Closing Excel and cleaning up...")
				closeExcel()
				time.Sleep(500 * time.Millisecond)
				maxRetries := 10
				for i := 0; i < maxRetries; i++ {
					err := os.Remove(tmpFile)
					if err == nil {
						fmt.Println("Temporary export file deleted.")
						return
					}

					fmt.Println("File still in use, retrying in 500ms...")
					time.Sleep(500 * time.Millisecond)
				}

				fmt.Println("Failed to delete temporary file after multiple attempts.")
			}()

			f := excelize.NewFile()

			sheet := "Sheet1"
			// write column headers
			f.SetCellValue(sheet, "A1", "ID")
			f.SetCellValue(sheet, "B1", "Date")
			f.SetCellValue(sheet, "C1", "Amount")
			f.SetCellValue(sheet, "D1", "Store")
			f.SetCellValue(sheet, "E1", "Category")

			largestIdWidth := 0
			largestDateWidth := 0
			largestCategoryWidth := 0
			largestAmountWidth := 0
			largestStoreWidth := 0
			// fill out the sheet with what we need
			for rowIndex, expense := range queriedExpenses {
				rowNumber := rowIndex + 2 // account for 0-based index and column headers

				idValue := expense.ID.Hex()
				dateValue := misc.FormatDateMMDDYYYY(expense.Date)
				f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNumber), idValue)
				f.SetCellValue(sheet, fmt.Sprintf("B%d", rowNumber), dateValue)
				f.SetCellValue(sheet, fmt.Sprintf("C%d", rowNumber), expense.Amount)
				f.SetCellValue(sheet, fmt.Sprintf("D%d", rowNumber), expense.Store)
				f.SetCellValue(sheet, fmt.Sprintf("E%d", rowNumber), expense.Category)

				idWidth := utf8.RuneCountInString(idValue) + 2
				dateWidth := utf8.RuneCountInString(dateValue) + 2
				amountWidth := utf8.RuneCountInString(strconv.FormatFloat(expense.Amount, 'f', 2, 64)) + 2
				storeWidth := utf8.RuneCountInString(expense.Store) + 2
				categoryWidth := utf8.RuneCountInString(expense.Category) + 2
				if idWidth > largestIdWidth {
					largestIdWidth = idWidth
				}
				if dateWidth > largestDateWidth {
					largestDateWidth = dateWidth
				}
				if amountWidth > largestAmountWidth {
					largestAmountWidth = amountWidth
				}
				if storeWidth > largestStoreWidth {
					largestStoreWidth = storeWidth
				}
				if categoryWidth > largestCategoryWidth {
					largestCategoryWidth = categoryWidth
				}
			}

			f.SetColWidth(sheet, "A", "A", float64(largestIdWidth))
			f.SetColWidth(sheet, "B", "B", float64(largestDateWidth))
			f.SetColWidth(sheet, "C", "C", float64(largestAmountWidth))
			f.SetColWidth(sheet, "D", "D", float64(largestStoreWidth))
			f.SetColWidth(sheet, "E", "E", float64(largestCategoryWidth))

			if err := f.SaveAs(tmpFile); err != nil {
				fmt.Println("Error saving file:", err)
			}

			openExcel(tmpFile)
		}
	},
}

func openExcel(file string) {
	var command *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		command = exec.Command("cmd", "/C", "start", "", file)
	case "darwin":
		command = exec.Command("open", "-a", "Microsoft Excel", file)
	default:
		command = exec.Command("xdg-open", file)
	}

	err := command.Start()
	if err != nil {
		fmt.Println("Error opening Excel:", err)
	}

	fmt.Println("Excel opened. Press Enter to close Excel and clean up...")
	fmt.Scanln()
}

func closeExcel() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("taskkill", "/IM", "EXCEL.EXE", "/F") // Force kill Excel on Windows
	case "darwin":
		cmd = exec.Command("pkill", "Microsoft Excel") // Close Excel on macOS
	default:
		cmd = exec.Command("pkill", "libreoffice") // Close LibreOffice on Linux (adjust if using another program)
	}

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error closing Excel:", err)
	} else {
		fmt.Println("Excel closed successfully.")
	}
}

func init() {
	summarizeExpenses.Flags().StringVarP(&filterMonth, "month", "m", "", "Specify the month to get a summary for (MM-YY)")
	summarizeExpenses.Flags().StringVarP(&filterCategory, "category", "c", "", "Specify the category to get a summary for")
	summarizeExpenses.Flags().StringVarP(&filterStore, "store", "s", "", "Specify the store in which your summary will apply to")
	summarizeExpenses.Flags().BoolVarP(&export, "export", "e", false, "Use if you wish to open the queried expenses in a spreadsheet format")

	RootCmd.AddCommand(summarizeExpenses)
}
