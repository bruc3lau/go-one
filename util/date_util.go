package util

import (
	"time"
)

// BetweenDays calculates the number of days between two dates
func BetweenDays(start, end string) (int64, error) {
	startDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return 0, err
	}
	endDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		return 0, err
	}
	return int64(endDate.Sub(startDate).Hours() / 24), nil
}

// BetweenDays2 calculates the number of days between two dates, considering current date as upper limit
func BetweenDays2(start, end string) (int64, error) {
	startDate, err := time.Parse("2006-01-02", start)
	if err != nil {
		return 0, err
	}
	endDate, err := time.Parse("2006-01-02", end)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	if endDate.After(now) {
		endDate = now
	}

	return int64(endDate.Sub(startDate).Hours()/24) + 1, nil
}

// StrToDate parses a string to time.Time
func StrToDate(str string) (time.Time, error) {
	if len(str) == 10 {
		return time.Parse("2006-01-02", str)
	}
	if len(str) == 19 {
		return time.Parse("2006-01-02 15:04:05", str)
	}
	return time.Time{}, &time.ParseError{Layout: "2006-01-02 or 2006-01-02 15:04:05", Value: str, LayoutElem: ""}
}
