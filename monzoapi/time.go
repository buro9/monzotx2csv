package monzoapi

import (
	"strings"
	"time"
)

// Time is a time.Time and is here to allow custom JSON unmarshalling so that
// we can read Monzo time format.
type Time struct {
	time.Time
}

// UnmarshalJSON is used by json.Unmarshal to parse a string into the Time
// struct using a custom format
func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		// null time, i.e. the Transaction.Settled field
		// use t.IsZero() to detect
		return nil
	}

	tt, err := time.Parse(time.RFC3339, s)
	if err != nil {
		// Monzo use a time format not listed at
		// https://golang.org/pkg/time/#pkg-constants
		tt, err = time.Parse("2006-01-02T15:04:05.000Z", s)
		if err != nil {
			return err
		}
	}

	*t = Time{tt}

	return nil
}
