package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05" // Correct layout for the given date format
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05" // example date "July 25, 2019 13:45:00"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05" // example date "Thursday, July 25, 2019 13:45:00"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}

	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05" // example date "7/25/2019 13:45:00"
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse date: %v", err))
	}

	return (fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute()))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now().UTC()
	currentYear := now.Year()

	// Create the anniversary date for the current year
	anniversary := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)

	return anniversary
}
