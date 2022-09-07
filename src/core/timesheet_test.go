package core

import (
	"testing"
	"time"
)

func Test__The_user_can_get_a_list_of_durations_per_bucket(t *testing.T) {
	now := time.Now()
	twoHours, _ := time.ParseDuration("2h")
	threeHours, _ := time.ParseDuration("3h")

	timesheet := NewTimesheet(
		[]TimesheetEntry{
			{now, now.Add(threeHours), "B", "bla"},
			{now, now.Add(twoHours), "B", "bla"},
			{now, now.Add(twoHours), "C", "bla"},
		},
	)

	bucketDurations := timesheet.GetDurationSumsPerBucket()

	if len(bucketDurations) != 2 {
		t.Errorf("Expected one bucket duration, but got %d", len(bucketDurations))
	}
	if bucketDurations[0] != "B: 05:00" {
		t.Errorf("Expected a duration of 05:00 for bucket B, but got %s", bucketDurations[0])
	}
	if bucketDurations[1] != "C: 02:00" {
		t.Errorf("Expected a duration of 02:00 for bucket C, but got %s", bucketDurations[1])
	}
}
