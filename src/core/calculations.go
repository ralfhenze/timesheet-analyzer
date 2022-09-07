package core

import (
	"fmt"
	"time"
)

func GetDurationSumsPerBucket(entries []TimesheetEntry) []string {
	var durationSums []string
	buckets := make(map[string][]TimesheetEntry)

	for _, entry := range entries {
		buckets[entry.Bucket] = append(buckets[entry.Bucket], entry)
	}

	for bucket, bucketEntries := range buckets {
		var duration time.Duration

		for _, bucketEntry := range bucketEntries {
			duration += bucketEntry.GetDuration()
		}

		hours := int(duration.Minutes()) / 60
		minutes := int(duration.Minutes()) % 60

		durationSums = append(durationSums, fmt.Sprintf("%s: %02d:%02d", bucket, hours, minutes))
	}

	return durationSums
}

func GetOverallDuration(entries []TimesheetEntry) string {

	var durationSum time.Duration

	for _, entry := range entries {
		durationSum += entry.GetDuration()
	}

	hours := int(durationSum.Minutes()) / 60
	minutes := int(durationSum.Minutes()) % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
