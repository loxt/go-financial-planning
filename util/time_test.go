package util

import "testing"

func TestStringToTime(t *testing.T) {
	var convertedTime = StringToTime("2020-10-03T14:21:10")

	if convertedTime.Year() != 2020 {
		t.Errorf("expects the year to be 2020")
	}

	if convertedTime.Month() != 10 {
		t.Errorf("expects the month to be 10")
	}
}
