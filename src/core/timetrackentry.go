package core

import "time"

type TimeTrackEntry struct {
	Start   time.Time
	End     time.Time
	Bucket  string
	Comment string
}

func (entry *TimeTrackEntry) GetDuration() time.Duration {
	return entry.End.Sub(entry.Start)
}
