package utils

import (
	"fmt"
	"strings"
)

func GetErrorIdxCheck(line string, Error []string) int {
	for _, err := range Error {
		idx := strings.LastIndex(line, err)
		fmt.Println(idx)
		if idx == -1 {
			return -1
		}
	}
	return 0
}
