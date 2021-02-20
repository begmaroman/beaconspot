package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLastDayOfMonth(t *testing.T) {
	testTable := []*struct {
		testName    string
		date        time.Time
		daysInMonth int
	}{
		{
			testName:    "January",
			date:        time.Date(2019, time.January, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 31,
		},
		{
			testName:    "February",
			date:        time.Date(2019, time.February, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 28,
		},
		{
			testName:    "February (leap)",
			date:        time.Date(2016, time.February, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 29,
		},
		{
			testName:    "March",
			date:        time.Date(2019, time.March, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 31,
		},
		{
			testName:    "April",
			date:        time.Date(2019, time.April, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 30,
		},
		{
			testName:    "May",
			date:        time.Date(2019, time.May, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 31,
		},
		{
			testName:    "June",
			date:        time.Date(2019, time.June, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 30,
		},
		{
			testName:    "July",
			date:        time.Date(2019, time.July, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 31,
		},
		{
			testName:    "August",
			date:        time.Date(2019, time.August, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 31,
		},
		{
			testName:    "September",
			date:        time.Date(2019, time.September, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 30,
		},
		{
			testName:    "October",
			date:        time.Date(2019, time.October, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 31,
		},
		{
			testName:    "November",
			date:        time.Date(2019, time.November, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 30,
		},
		{
			testName:    "December",
			date:        time.Date(2019, time.December, 5, 0, 0, 0, 0, time.UTC),
			daysInMonth: 31,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.testName, func(t *testing.T) {
			daysInMonth := LastDayOfMonth(tt.date)
			require.Equal(t, tt.daysInMonth, daysInMonth)
		})
	}
}
