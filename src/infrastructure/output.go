package infrastructure

import (
	"fmt"
	"time"
)

func GetDurationStringHhMm(duration time.Duration) string {
	hours := int(duration.Minutes()) / 60
	minutes := int(duration.Minutes()) % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
