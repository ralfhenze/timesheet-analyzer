package infrastructure

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"
	"timesheet-analyzer/core"
)

func GetTimesheetFromCsvString(csvString string) core.Timesheet {
	buffer := bytes.NewReader([]byte(csvString))

	csvReader := csv.NewReader(buffer)
	csvReader.Comma = '\t'
	csvLines, _ := csvReader.ReadAll()

	var entries []core.TimesheetEntry

	for i, line := range csvLines {
		var entry core.TimesheetEntry

		entry.Start, _ = time.Parse("2006-01-02 15:04:05", line[0]+" "+line[1]+":00")
		entry.End, _ = time.Parse("2006-01-02 15:04:05", line[0]+" "+line[2]+":00")

		entry.Bucket = line[3]
		entry.Comment = line[4]

		fmt.Printf("%d: %s\n", i, entry)

		entries = append(entries, entry)
	}

	return core.NewTimesheet(entries)
}
