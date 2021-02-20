package timex

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Time is just like time.Time, except it maintains UTC and marshals
// to second precision only
//
// Use .Time to access the underlying time.Time.
type Time struct {
	time.Time
}

const rfc3339MillisecondsFormat = "2006-01-02T15:04:05.000Z07:00"

// NewTime creates a Time instance from a time.Time instance (standardized to UTC).
func NewTime(t time.Time) Time {
	return Time{t.UTC()}
}

// NewTimeFromString parses a Time instance from a string using RFC
// 3339 parsing.
func NewTimeFromString(s string) (Time, error) {
	t, err := time.Parse(rfc3339MillisecondsFormat, s)
	return NewTime(t), err
}

// Now returns the current time (standardized to UTC)
func Now() Time {
	return NewTime(time.Now())
}

// MarshalJSON serializes a Time instance as the underlying time.Time instance.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.RFC3339Format())), nil
}

// UnmarshalJSON deserializes a Time instance from the JSON string representation.
func (t *Time) UnmarshalJSON(data []byte) error {
	decoded, err := NewTimeFromString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}

	*t = decoded
	return nil
}

// RFC3339Format returns the RFC 3339 string representation of the time,
// set at UTC (terminated in a Z), in the resolution of milliseconds
func (t Time) RFC3339Format() string {
	return t.Format(rfc3339MillisecondsFormat)
}

// Value implements serialization for database/sql
func (t Time) Value() (driver.Value, error) {
	return t.Time, nil
}

// Scan implements deserialization for database/sql
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		return errors.New("unable scan empty value into timex.Time")
	}

	value, err := driver.String.ConvertValue(value)
	if err != nil {
		return fmt.Errorf("unable to convert value into string: %v", value)
	}

	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("unable to convert value into string: %v", value)
	}

	inner, err := time.Parse("2006-01-02 15:04:05 -0700 MST", s)
	if err != nil {
		return errors.Wrapf(err, "unable to parse string into timex.Time: %v", s)
	}

	t.Time = inner.UTC()
	return nil
}

// Iso3339CleanTime converts the given time to ISO 3339 format
func Iso3339CleanTime(t time.Time) string {
	return t.UTC().Format("2006-01-02T15:04:05.000000Z")
}
