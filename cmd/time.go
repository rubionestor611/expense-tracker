/*
Copyright © 2025 nestordrubio9@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"example.com/nestor-expense-tracker/misc"
	"github.com/spf13/cobra"
)

// timezoneCmd represents the timezone command
var timezoneCmd = &cobra.Command{
	Use:   "time",
	Short: "An interesting little side tool to start with Cobra CLI building",
	Long:  `An interesting little side tool to start with Cobra CLI building. It allows me to see tiems in timezones or locally if I don't define a timezone`,
	Run: func(cmd *cobra.Command, args []string) {
		var timezone string

		if len(args) > 0 {
			timezone = args[0]
		} else {
			localTimezone, _ := time.Now().Zone()

			timezone = localTimezone
		}

		currentTime, err := misc.GetTimeInTimezone(timezone)
		if err != nil {
			log.Fatalf("The timezone string '%s' is invalid", timezone)
		}
		fmt.Println(currentTime)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timezoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timezoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
