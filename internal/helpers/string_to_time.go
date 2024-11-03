package helpers

import "time"

func StringToTime(timeStr *string, layout string) *time.Time {
	if timeStr == nil {
		return nil
	}

	parsedTime, err := time.Parse(layout, *timeStr)
	if err != nil {
		return nil
	}

	return &parsedTime
}
