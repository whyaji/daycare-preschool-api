package utils

import (
	"fmt"
	"time"
)

// convert string to time.Time
func ParseDateStringToTime(dateStr string) (*time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format. use YYYY-MM-DD")
	}
	return &date, nil
}

func ParseDateTimeStringToTime(dateTimeStr string) (*time.Time, error) {
	dateTime, err := time.Parse("2006-01-02 15:04:05", dateTimeStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date time format. use YYYY-MM-DD HH:mm:ss")
	}
	return &dateTime, nil
}
