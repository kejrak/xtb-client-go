package xtb

import (
	"fmt"
	"strconv"
	"time"
)

// Time is a wrapper around time.Time to handle custom JSON serialization/deserialization.
type Time time.Time

// ToTime converts the custom Time type to the standard time.Time type.
func (t *Time) ToTime() time.Time {
	return time.Time(*t)
}

// UnmarshalJSON customizes the unmarshalling of Time from JSON.
// It converts a JSON timestamp (milliseconds since Unix epoch) to a Time object.
func (t *Time) UnmarshalJSON(text []byte) error {
	timestamp, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}

	*t = Time(time.UnixMilli(timestamp))
	return nil
}

// MarshalJSON customizes the marshalling of Time to JSON.
// It converts a Time object to its JSON representation (timestamp in milliseconds).
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).UnixMilli(), 10)), nil
}
