package infrastructure

import (
	"testing"
	"time"
)

func Test__Durations_are_shown_in_HH_MM_format(t *testing.T) {
	oneHour, _ := time.ParseDuration("1h")

	durationString := GetDurationStringHhMm(oneHour)

	if durationString != "01:00" {
		t.Errorf("Expected duration string \"01:00\", but was \"%s\"", durationString)
	}
}
