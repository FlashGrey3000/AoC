package main

import (
	"strings"

	utils "github.com/FlashGrey3000/AoC/utils"
)

func main() {
	lines := utils.readLines("input.txt")

	for i:=0; i<len(lines); i++ {
		line := strings.Split(lines[i], ": ")
		left := utils.Atoi(line[0])
		rrr := strings.Split(line[1])
		right := utils.AvectoIvec(rrr)
		
	}

}

func isCalculationAMatch(expectedSum, sum int, input []int) bool {
	if len(input) == 0 {
		return sum == expectedSum
	}

	if sum > expectedSum {
		return false
	}

	if isCalculationAMatch(expectedSum, calculate(sum, input[0], '+'), input[1:]) {
		return true
	}

	return isCalculationAMatch(expectedSum, calculate(sum, input[0], '*'), input[1:])
}

func calculate(a, b int, operation byte) int {
	calculation := 0
	switch operation {
	case '+':
		calculation = a + b
	case '*':
		calculation = a * b
	}
	return calculation
}