package core

import (
	"time"
)

type Timesheet struct {
	entries []TimesheetEntry
}

func NewTimesheet(entries []TimesheetEntry) Timesheet {
	return Timesheet{
		entries: entries,
	}
}

func (ts *Timesheet) GetDurationsPerBucket() map[string]time.Duration {
	durationSums := make(map[string]time.Duration)
	buckets := make(map[string][]TimesheetEntry)

	for _, entry := range ts.entries {
		buckets[entry.Bucket] = append(buckets[entry.Bucket], entry)
	}

	for bucket, bucketEntries := range buckets {
		var duration time.Duration

		for _, bucketEntry := range bucketEntries {
			duration += bucketEntry.GetDuration()
		}

		durationSums[bucket] = duration
	}

	return durationSums
}

func (ts *Timesheet) GetOverallDuration() time.Duration {

	var overallDuration time.Duration

	for _, entry := range ts.entries {
		overallDuration += entry.GetDuration()
	}

	return overallDuration
}
