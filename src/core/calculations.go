package core

import (
	"fmt"
	"time"
)

func GetDurationSumsPerBucket(entries []TimeTrackEntry) []string {
	var durationSums []string
	buckets := make(map[string][]TimeTrackEntry)

	for _, entry := range entries {
		buckets[entry.Bucket] = append(buckets[entry.Bucket], entry)
	}

	for bucket, bucketEntries := range buckets {
		var duration time.Duration

		for _, bucketEntry := range bucketEntries {
			duration += bucketEntry.Duration
		}

		hours := int(duration.Minutes()) / 60
		minutes := int(duration.Minutes()) % 60

		durationSums = append(durationSums, fmt.Sprintf("%s: %02d:%02d", bucket, hours, minutes))
	}

	return durationSums
}

func GetOverallDuration(entries []TimeTrackEntry) string {

	var durationSum time.Duration

	for _, entry := range entries {
		durationSum += entry.Duration
	}

	hours := int(durationSum.Minutes()) / 60
	minutes := int(durationSum.Minutes()) % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
