package utils

import "time"

const (
	layoutISO = "2006-01-02"
)

func Date(date string) (time.Time, error) {
	return time.Parse(layoutISO, date)
}
