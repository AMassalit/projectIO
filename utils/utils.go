package utils

import (
	"fmt"
	"time"
)

func FormatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%d ч %d мин %d сек", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("%d мин %d сек", minutes, seconds)
	}
	return fmt.Sprintf("%d сек", seconds)
}
