package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"timesheet-analyzer/infrastructure"
)

func main() {
	stdinBytes, _ := ioutil.ReadAll(os.Stdin)

	timesheet := infrastructure.GetTimesheetFromCsvString(string(stdinBytes))

	fmt.Println("\n--------------- Durations per Bucket ---------------")

	durations := timesheet.GetDurationsPerBucket()

	for bucket, duration := range durations {
		fmt.Println(bucket + ": " + infrastructure.GetDurationStringHhMm(duration))
	}

	fmt.Println("\n--------------- Overall Duration ---------------")
	overallDuration := timesheet.GetOverallDuration()

	fmt.Println(infrastructure.GetDurationStringHhMm(overallDuration))
}
