package main

import (
	"fmt"
	"strconv"
	"strings"
	"os"
	"regexp"
)

func main() {
	data,err := os.ReadFile("input.txt")
	if err!=nil {
		fmt.Println(err)
	}

	// peta := []byte(`xmul(2,4)%&mul[3,7]!@^don't()?_mul(5,5)+mul(32,64]don't()_then_do()(mul(11,8)mul(8,5))`)
	
	sum := 0
	re := regexp.MustCompile(`(mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\))`)
	
	decodedInstructions := re.FindAll(data, -1)
	enabled := true
	mul := 0

	for i, j:= range decodedInstructions {
		if strings.Compare(string(j), "don't()") == 0 {
			fmt.Println(i,string(j), "  sum:", sum, "   enabled: ", enabled)
			enabled = false
			mul = 0
		} else if strings.Compare(string(j), "do()") == 0 {
			fmt.Println(i,string(j), "  sum:", sum, "   enabled: ", enabled)
			enabled = true
			mul = 0
		} else {
			fmt.Println(i,string(j), "  sum:", sum, "   enabled: ", enabled)
			part := strings.Split(string(j), ",")
			//fmt.Println(part[0], part[1])
			part1,err := strconv.Atoi(part[0][4:])
			if err!=nil {
				fmt.Println(err)
			}
			part2,err := strconv.Atoi(part[1][:len(part[1])-1])
			if err!=nil {
				fmt.Println(err)
			}
			mul = part1 * part2
		}

		if enabled {
			sum += mul
		}
	}

	fmt.Println("sum = ", sum)

}