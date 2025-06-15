package utils

import "fmt"

func CmtOutput(m map[string]any) string {
	var value string

	for k, v := range m {
		str := fmt.Sprintln(k, ":", v)
		value += str
	}
	return value
}
