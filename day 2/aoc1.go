package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(x string) (bool) {
	levels := strings.Split(x, " ")

	incinit := true

	for i:=0; i<len(levels) - 1; i++ {
		a, err := strconv.Atoi(strings.TrimSpace(levels[i]))
		if err!=nil {
			fmt.Println(err)
		}
		b, err := strconv.Atoi(strings.TrimSpace(levels[i+1]))
		if err!=nil {
			fmt.Println(err)
		}

		if a-b>0 && i==0 {
			incinit = false
		}

		inc := true
		if a-b>0 {
			inc = false
		}

		if inc {
			if b-a < 1 || b-a > 3 || inc != incinit{
				return false
			}
		} else {
			if a-b < 1 || a-b > 3 || inc != incinit{
				return false
			}
		}
		
	}

	return true
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err!=nil {
		fmt.Println(err)
		return
	}

	reports := strings.Split(string(data), "\n")

	safeReports := 0

	for i:=0; i<1000;i++ {
		
		pp:=isSafe(reports[i])
		if pp {
			safeReports++
		} else {
			safety:=dampenProblem(reports[i])
			if safety {
				safeReports++
			}
		}
		fmt.Printf("report %d: %s : safe? : %t\n", i, strings.TrimSpace(reports[i]), pp)
	}

	fmt.Println("No. of safe reports: ", safeReports)
}