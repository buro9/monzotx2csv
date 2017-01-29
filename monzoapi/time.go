package monzoapi

import (
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		// nullable time, i.e. the Transaction.Settled field
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
