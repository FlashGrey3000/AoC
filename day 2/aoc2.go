package main

import (
	"strings"
)

func dampenProblem(x string) bool {

	levels := strings.Split(x, " ")
	for i := 0; i < len(levels); i++ {
		var result []string

		for j, level := range levels {
			if j!=i {
				result = append(result, strings.TrimSpace(level))
			}
		}

		resultf := strings.Join(result, " ")

		safety := isSafe(resultf)

		if safety == true {
			return true
		}
	}
	return false
}