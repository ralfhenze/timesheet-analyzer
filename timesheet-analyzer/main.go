package main

import (
	"fmt"
	"timesheet-analyzer/infrastructure"
)

func main() {

	timesheet := infrastructure.GetTimesheetFromCsvFile(
		"../examples/timesheet1.csv",
	)

	fmt.Println("\n--------------- Durations per Bucket ---------------")

	durations := timesheet.GetDurationsPerBucket()

	for bucket, duration := range durations {
		fmt.Println(bucket + ": " + infrastructure.GetDurationStringHhMm(duration))
	}

	fmt.Println("\n--------------- Overall Duration ---------------")
	overallDuration := timesheet.GetOverallDuration()

	fmt.Println(infrastructure.GetDurationStringHhMm(overallDuration))
}
