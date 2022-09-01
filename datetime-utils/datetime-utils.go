package datetime_utils

import (
	"time"
)

// ConvertStringToTime returns a time.Time struct that
// is parsed from the given dateString in format "YYYY-MM-DD".
// Will panic if the dateString is in an unexpected format.
func ConvertStringToTime(dateString string) time.Time {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		panic(err)
	}
	return date
}

// FindDayOfWeek returns a string with the day of the week
// for a given time.Time
func FindDayOfWeek(datetime time.Time) string {
	return datetime.Weekday().String()
}

// ConvertUnixTimeToRFC3339 returns a string representation
// in RFC3339 format of the time converted from the given unixTime.
func ConvertUnixTimeToRFC3339(unixTime int64) string {
	timezone, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}
	convertedTime := time.Unix(unixTime, 0)
	return convertedTime.In(timezone).Format(time.RFC3339)
}
