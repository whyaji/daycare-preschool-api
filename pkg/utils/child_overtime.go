package utils

import "time"

func CalculateChildMorningOvertime(inputTime time.Time) int {
	// cutoff 07:45:00
	cutoff := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 7, 45, 0, 0, inputTime.Location())

	if inputTime.Before(cutoff) {
		diff := cutoff.Sub(inputTime)
		totalMinutes := int(diff.Minutes())
		overtime := (totalMinutes + 14) / 15
		return overtime
	}
	return 0
}

func CalculateChildEveningOvertime(inputTime time.Time) int {
	// cutoff 16:15:00
	cutoff := time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day(), 16, 15, 0, 0, inputTime.Location())

	if inputTime.After(cutoff) {
		diff := inputTime.Sub(cutoff)
		totalMinutes := int(diff.Minutes())
		overtime := (totalMinutes + 14) / 15
		return overtime
	}
	return 0
}
