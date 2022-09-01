package datetime_utils

import (
	"regexp"
	"testing"
	"time"
)

// TestConvertStringToTimePanic verifies that the function
// panics if an unexpected format is passed in
func TestConvertStringToTimePanic(t *testing.T) {
	defer func() { recover() }()
	ConvertStringToTime("01-02-2022")
	t.Errorf("TestConvertStringToTimePanic should have panicked!")
}

// TestConvertStringToTime is essentially a smoke test wrapper around
// time.Parse() to verify that we're receiving the output we expect.
// We don't need to create unit tests that verify time.Parse() itself works.
func TestConvertStringToTime(t *testing.T) {
	response := ConvertStringToTime("2022-01-01")
	expected := time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)
	if response != expected {
		t.Errorf("expected result to be %v, received %v", expected, response)
	}
}

// TestFindDayOfWeek is essentially a smoke test of the wrapper around
// time.Weekday().String() to verify that we're receiving the output we expect.
// We don't need to create unit tests that verify time.Weekday().String() works.
func TestFindDayOfWeek(t *testing.T) {
	datetime := time.Now()
	expected := regexp.MustCompile("day")
	response := FindDayOfWeek(datetime)
	if !expected.MatchString(response) {
		t.Fatalf(`FindDayOfWeek(datetime) = %q, expected to end in -%#q`, response, expected)
	}
}

// TestConvertUnixTimeToRFC3339 Verifies that unix time is converted
// to a string in RFC3339 format using EST timezone
func TestConvertUnixTimeToRFC3339(t *testing.T) {
	unixEpoch := 0
	expected := regexp.MustCompile("1969-12-31T19:00:00-05:00")
	response := ConvertUnixTimeToRFC3339(int64(unixEpoch))
	if !expected.MatchString(response) {
		t.Fatalf(`TestConvertUnixTimeToRFC3339 = %q, expected to convert to EST timestamp %q`, response, expected)
	}
}
