package main

import (
	"fmt"
	"strings"

	utils "github.com/FlashGrey3000/AoC/utils"
)

func main() {
	lines := utils.ReadLines("input.txt")

	a1,a2:=0,0

	for i:=0; i<len(lines); i++ {
		line := strings.Split(lines[i], ": ")
		left := utils.Atoi(line[0])
		rrr := strings.Split(line[1], " ")
		right := utils.AvectoIvec(rrr)

		if isCalculationAMatch(left, 0, right, false) {
			a1+=left
		}

		if isCalculationAMatch(left, 0, right, true) {
			a2+=left
		}
		
	}

	fmt.Printf("answer 1: %d, answer 2: %d\n", a1, a2)

}

func isCalculationAMatch(expectedSum, sum int, input []int, isPart2 bool) bool {
	if len(input) == 0 {
		return sum == expectedSum
	}

	if sum > expectedSum {
		return false
	}

	if isCalculationAMatch(expectedSum, calculate(sum, input[0], '+'), input[1:], isPart2) {
		return true
	}

	if isPart2 && isCalculationAMatch(expectedSum, calculate(sum, input[0], '|'), input[1:], isPart2) {
		return true
	}
	return isCalculationAMatch(expectedSum, calculate(sum, input[0], '*'), input[1:], isPart2)
}

func calculate(a, b int, operation byte) int {
	calculation := 0
	switch operation {
	case '+':
		calculation = a + b
	case '*':
		calculation = a * b
	case '|':
		mul, q := 10, 10
		for q != 0 {
			q = b / mul
			if q > 0 {
				mul *= 10
			}
		}
		calculation = (a * mul) + b
	}
	return calculation
}