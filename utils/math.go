package utils

import "fmt"

func ConvertTimestamp(time float64) string {
	totalSecs := int64(time)
	hours := totalSecs / 3600
	minutes := (totalSecs % 3600) / 60
	seconds := totalSecs % 60

	var timeString string
	if hours != 0 {
		timeString = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	} else {
		timeString = fmt.Sprintf("%02d:%02d", minutes, seconds)
	}

	return timeString
}
