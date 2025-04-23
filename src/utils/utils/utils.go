package utils

import (
	"strings"
)

func GetErrorIdxCheck(line string, Error []string, needErrors bool) int {
	for _, err := range Error {
		idx := strings.LastIndex(line, err)
		if idx == -1 && !needErrors {
			return -1

		}
	}
	return 0
}
