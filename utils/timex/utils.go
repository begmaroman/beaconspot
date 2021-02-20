package timex

import (
	"time"
)

// LastDayOfMonth returns the last day of the month of the given date.
func LastDayOfMonth(dateInMonth time.Time) int {
	return time.Date(dateInMonth.Year(), dateInMonth.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
