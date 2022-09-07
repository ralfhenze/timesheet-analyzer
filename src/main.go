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

	durationSums := timesheet.GetDurationSumsPerBucket()

	for _, sum := range durationSums {
		fmt.Println(sum)
	}

	fmt.Println("\n--------------- Overall Duration ---------------")

	fmt.Println(timesheet.GetOverallDuration())
}
