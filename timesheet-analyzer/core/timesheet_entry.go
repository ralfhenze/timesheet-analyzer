package core

import "time"

type TimesheetEntry struct {
	Start   time.Time
	End     time.Time
	Bucket  string
	Comment string
}

func (entry *TimesheetEntry) GetDuration() time.Duration {
	return entry.End.Sub(entry.Start)
}
