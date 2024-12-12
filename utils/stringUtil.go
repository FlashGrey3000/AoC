package utils

import (
	"os"
	"strings"
)

func readLines(filename string) []string {
	file,err := os.ReadFile(filename)
	if err != nil {
		return []string{"file not found"}
	}

	return strings.Split(string(file), "\r\n")
}