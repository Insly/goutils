package goutils

import "time"

func CalculateBusinessDaysBetweenDates(from time.Time, to time.Time) int {
	days := 0
	for {
		if from.After(to) || from.Equal(to) {
			return days
		}
		if from.Weekday() != time.Saturday && from.Weekday() != time.Sunday {
			days++
		}
		from = from.Add(time.Hour * 24)
	}
}
