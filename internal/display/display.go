package display

import (
	"fmt"
	"time"
)

// PrintMap Helps print out a YAML map in a more human-readable way.
func PrintMap(m map[string]interface{}, prefix string) {
	for k, v := range m {
		if m, ok := v.(map[string]interface{}); ok {
			fmt.Printf("%s%s:\n", prefix, k)
			PrintMap(m, prefix+"  ")
		} else {
			fmt.Printf("%s%s: %v\n", prefix, k, v)
		}
	}
}

// FormatTime to change the time format to something more readable
func FormatTime(timeStr string) string {
	layout := "2006-01-02T15:04"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return ""
	}

	// Prepend a 0 if the time is =6 to keep things uniform
	kitchenFormat := parsedTime.Format(time.Kitchen)
	if len(kitchenFormat) == 6 {
		return "0" + kitchenFormat
	} else {
		return kitchenFormat
	}
}

// FormatDate by parsing the timeStr and using the time.DateOnly format
func FormatDate(timeStr string) string {
	layout := "2006-01-02T15:04"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return ""
	}

	return parsedTime.Format(time.DateOnly)
}

func FormatPercentage(percentage float64) string {
	if percentage < 10 {
		return fmt.Sprintf("0%v%%", percentage)
	} else {
		return fmt.Sprintf("%v%%", percentage)
	}
}
