package display

import (
	"fmt"
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
