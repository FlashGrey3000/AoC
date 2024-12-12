package utils

import (
	"os"
	"strconv"
	"strings"
)

func readLines(filename string) []string {
	file,err := os.ReadFile(filename)
	if err != nil {
		return []string{"file not found"}
	}

	return strings.Split(string(file), "\r\n")
}

func Atoi(a string) (int) {
	x,err := strconv.Atoi(a)
	if err != nil {
		return -99999999
	}
	return x
}

func AvectoIvec(aVec string) ([]int) {
	iVec := make([]int, len(aVec))
	for i,a := range aVec {
		iVec[i] = Atoi(string(a))
	}
	return iVec
}