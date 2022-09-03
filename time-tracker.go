package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type TimeTrackEntry struct {
	Start    time.Time
	End      time.Time
	Duration time.Duration
	Bucket   string
	Comment  string
}

func main() {
	csvFile, _ := os.Open("times.csv")
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = '\t'
	csvLines, _ := csvReader.ReadAll()

	var entries []TimeTrackEntry

	for i, line := range csvLines {
		var entry TimeTrackEntry

		entry.Start, _ = time.Parse("2006-01-02 15:04:05", line[0]+" "+line[1]+":00")
		entry.End, _ = time.Parse("2006-01-02 15:04:05", line[0]+" "+line[2]+":00")

		entry.Duration = entry.End.Sub(entry.Start)

		entry.Bucket = line[3]
		entry.Comment = line[4]

		fmt.Printf("%d: %s\n", i, entry)

		entries = append(entries, entry)
	}

	fmt.Println("\n-------------------------------------------------------------------------\n")

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

		fmt.Printf("%s: %02d:%02d\n", bucket, hours, minutes)
	}

	fmt.Println("\n-------------------------------------------------------------------------\n")

	var durationSum time.Duration

	for _, entry := range entries {
		durationSum += entry.Duration
	}

	hours := int(durationSum.Minutes()) / 60
	minutes := int(durationSum.Minutes()) % 60

	fmt.Printf("Overall: %02d:%02d\n", hours, minutes)
}
