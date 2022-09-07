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

	bucketDurations := timesheet.GetDurationsPerBucket()

	if len(bucketDurations) != 2 {
		t.Errorf("Expected two bucket durations, but got %d", len(bucketDurations))
	}
	fiveHours, _ := time.ParseDuration("5h")
	if bucketDurations["B"] != fiveHours {
		t.Errorf("Expected a duration of %s for bucket B, but got %s", fiveHours, bucketDurations["B"])
	}
	if bucketDurations["C"] != twoHours {
		t.Errorf("Expected a duration of %s for bucket C, but got %s", twoHours, bucketDurations["C"])
	}
}
