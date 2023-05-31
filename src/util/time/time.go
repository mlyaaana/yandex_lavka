package time

import (
	"time"
)

const (
	Day = 24 * time.Hour

	HourMinuteLayout = "15:04"
	DateLayout       = "2006-01-02"
)

func Max(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

func NowDate() time.Time {
	return time.Now().UTC().Truncate(Day)
}
