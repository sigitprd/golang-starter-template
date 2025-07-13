package utils

import (
	"fmt"
	"time"
	_ "time/tzdata"

	"github.com/pkg/errors"
)

const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

func Now() time.Time {
	return time.Now().UTC()
}

func NowJkt() time.Time {
	return time.Now().In(LocJkt())
}

func NowDateJkt() string {
	return NowJkt().Format(DateFormat)
}

func ParseDateString(oldFormat string, newFormat string, date string) (string, error) {
	dateTime, err := time.Parse(oldFormat, date)
	if err != nil {
		return "", errors.Wrap(err, "cannot parse date string")
	}
	return dateTime.Format(newFormat), nil
}

func LocJkt() *time.Location {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	return loc
}

// MonthsCountUntil calculates the months between now
// and the until time.Time value passed
func MonthsCountUntil(until time.Time) int {
	now := time.Now()
	months := 0
	month := until.Month()
	for now.Before(until) {
		now = now.Add(time.Hour * 24)
		nextMonth := now.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}

	if months == 0 {
		months = 1
	}

	return months
}

func ConvertDateJKTToUTCRange(startDateStr, endDateStr string) (string, string) {
	loc := LocJkt()
	// Parse startDate dan endDate di zona waktu Jakarta
	startDate, _ := time.ParseInLocation(DateFormat, startDateStr, loc)
	endDate, _ := time.ParseInLocation(DateFormat, endDateStr, loc)

	// Set waktu startDate ke 00:00:00 dan endDate ke 23:59:59, lalu convert ke UTC
	startDateTimeStr := startDate.Add(time.Hour * 0).UTC().Format(DateTimeFormat)
	endDateTimeStr := endDate.Add(time.Hour*23 + time.Minute*59 + time.Second*59).UTC().Format(DateTimeFormat)

	return startDateTimeStr, endDateTimeStr
}

func ConvertTimeUtcToJkt(at time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return at.In(loc)
}

// StringToConvertTime layout "2006-Jan-02" converts a string time to a time.Time pointer with timezone +07.
func StringToConvertTime(dateStr *string) (*time.Time, error) {
	if dateStr == nil || *dateStr == "" {
		return nil, nil
	}

	// Parse the string with the custom layout
	parsedTime, err := time.Parse(DateFormat, *dateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse time: %v", err)
	}

	// Convert the parsed time to the specified timezone
	convertedTime := parsedTime.In(LocJkt())

	return &convertedTime, nil
}
