package main

import (
	"fmt"
	"timesheet-analyzer/core"
	"timesheet-analyzer/infrastructure"
)

func main() {

	entries := infrastructure.GetTimesheetEntriesFromCsvFile(
		"../examples/timesheet1.csv",
	)

	fmt.Println("\n--------------- Durations per Bucket ---------------")

	durationSums := core.GetDurationSumsPerBucket(entries)

	for _, sum := range durationSums {
		fmt.Println(sum)
	}

	fmt.Println("\n--------------- Overall Duration ---------------")

	fmt.Println(core.GetOverallDuration(entries))
}
