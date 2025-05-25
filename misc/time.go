/*
Copyright © 2025 nestordrubio9@gmail.com
*/
package misc

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func GetTimeInTimezone(timezone string) (*time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}
	currentTime := time.Now().In(location)
	return &currentTime, nil
}

func ISOFormat(timeVal time.Time) string {
	return timeVal.Format(time.RFC3339)
}

func IsValidMMYY(dateStr string) bool {
	re := regexp.MustCompile(`^(0[1-9]|1[0-2])-\d{2}$`)
	return re.MatchString(dateStr)
}

func IsValidYYYY(yearStr string) bool {
	if len(yearStr) != 4 {
		return false
	}

	_, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}

	return true
}

func GetMonthRange(dateStr string) (time.Time, time.Time, error) {
	if !IsValidMMYY(dateStr) {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid date format, expected MM-YY")
	}

	// extract month and year
	parts := regexp.MustCompile(`-`).Split(dateStr, 2)
	month := parts[0]
	year := fmt.Sprintf("20%s", parts[1])

	startDate, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-01", year, month))
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	// get end of same month (add 1 month and remove one day)
	endDate := startDate.AddDate(0, 1, -1)

	return startDate, endDate, nil
}

func FormatDateMMDDYYYY(date time.Time) string {
	return date.Format("01-02-2006")
}
