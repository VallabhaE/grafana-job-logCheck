package utils

import "strings"

func GetErrorIdxCheck(line string, Error []string) int {
	for _, err := range Error {
		idx := strings.LastIndex(line, err)
		if idx == -1 {
			return -1
		}
	}
	return 0
}
