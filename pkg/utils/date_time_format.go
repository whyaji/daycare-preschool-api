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
