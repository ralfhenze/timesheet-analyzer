package core

import "time"

type TimeTrackEntry struct {
	Start    time.Time
	End      time.Time
	Duration time.Duration
	Bucket   string
	Comment  string
}
