package main

import (
	"fmt"
	"timesheet-analyzer/core"
	"timesheet-analyzer/infrastructure"
)

func main() {

	entries := infrastructure.GetTimeTrackEntriesFromCsvFile(
		"/home/ralf/Business/2022/Leistungsnachweise/Stundenzettel/Nextbike Stundenzettel 2022-08.csv",
	)

	fmt.Println("\n--------------- Durations per Bucket ---------------")

	durationSums := core.GetDurationSumsPerBucket(entries)

	for _, sum := range durationSums {
		fmt.Println(sum)
	}

	fmt.Println("\n--------------- Overall Duration ---------------")

	fmt.Println(core.GetOverallDuration(entries))
}
