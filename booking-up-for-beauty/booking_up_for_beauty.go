package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	today := time.Now()
	compareDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	isPast := compareDate.Before(today)
	return isPast
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	compareDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	compareHour := compareDate.Hour()
	if compareHour >= 12 && compareHour < 18{
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	returnDate := Schedule(date)
	return "You have an appointment on " + returnDate.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
