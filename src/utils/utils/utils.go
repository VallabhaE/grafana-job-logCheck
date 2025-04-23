package utils

import (
	"strings"
)

var Counter map[string]int

func GtErrorIdxCheck(line string, Error []string, needErrors bool) int {
	if Counter == nil {
		Counter = make(map[string]int)
	}
	if needErrors {
		for _, err := range Error {
			idx := strings.LastIndex(line, err)
			if idx != -1 {
				Counter[err]++
				return idx
			}
		}
		return -1
	} else {
		for _, err := range Error {
			idx := strings.LastIndex(line, err)
			if idx == -1 {
				continue
			}else{
				return -1
			}
		}
		return 0
	}

}
