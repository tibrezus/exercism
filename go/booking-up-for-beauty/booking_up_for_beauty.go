package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
// => 2019-07-25 13:45:00 +0000 UTC
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	parseDate, erro := time.Parse(layout, date)
	if erro != nil {
		return false
	}

	currentTime := time.Now()

	return parseDate.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	parseDate, erro := time.Parse(layout, date)
	if erro != nil {
		return false
	}

	hour := parseDate.Hour()

	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"

	parseDate, erro := time.Parse(layout, date)
	if erro != nil {
		return ""
	}

	return "You have an appointment on " + parseDate.Weekday().String() + ", " + parseDate.Month().String() + " " + parseDate.Format("2") + ", " + parseDate.Format("2006") + ", at " + parseDate.Format("15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
