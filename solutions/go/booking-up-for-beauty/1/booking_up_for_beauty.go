package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	sTime, _ := time.Parse(
		"1/2/2006 15:04:05", date)
	return sTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}

	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}

	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	// Format: "Thursday, July 25, 2019, at 13:45"
	return fmt.Sprintf(
		"You have an appointment on %s, %s %d, %d, at %02d:%02d.",
		t.Weekday().String(),
		t.Month().String(),
		t.Day(),
		t.Year(),
		t.Hour(),
		t.Minute(),
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now().UTC()
	currentYear := now.Year()

	// Anniversary on September 15 of the current year
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
