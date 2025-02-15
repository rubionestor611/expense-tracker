/*
Copyright © 2025 nestordrubio9@gmail.com
*/
package misc

import "time"

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
