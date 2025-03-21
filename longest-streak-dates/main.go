package main

import (
	"fmt"
	"time"
)

// TASK: Create a function that takes a list of date dictionaries and return the "longest streak" (i.e. longest number of consecutive days in a row).
func main() {
	longest := longestStreak([]string{
		"2019-09-18",
		"2019-09-19",
		"2019-09-20",
		"2019-09-26",
		"2019-09-27",
		"2019-09-30",
	})

	fmt.Println(longest)
}

func longestStreak(dates []string) int {
	if len(dates) == 0 {
		return 0
	}
	var longest, current int
	var prev time.Time
	for _, dt := range dates {
		d, err := convertStringToDate(dt)
		if err != nil {
			panic("invalid date")
		}
		if !isNextDay(d, prev) {
			current = 0
		}
		current++
		if current > longest {
			longest = current
		}
		prev = d
	}
	return longest
}

func convertStringToDate(dateStr string) (time.Time, error) {
	const layout = "2006-01-02" // Go's reference date for formatting and parsing
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

// isNextDay checks if two dates are next to each other separated by a day.
func isNextDay(date1, date2 time.Time) bool {
	// First, ensure both dates have the same location as daylight saving can affect comparisons.
	if date1.Location() != date2.Location() {
		return false
	}

	// Calculate the difference between the two dates in hours.
	diff := date2.Sub(date1).Hours()

	// If the difference is 24 hours, they are next to each other separated by a day.
	return diff == 24 || diff == -24
}
