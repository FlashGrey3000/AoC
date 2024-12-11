package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pagesBefore := make(map[int][]int, 0)
	
	data,err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	text := strings.Split(string(data), "$")

	rules := strings.Split(text[0], "\n")
	
	for i := 0; i < len(rules); i++ {
		rule := strings.Split(rules[i], "|")
		left,err := strconv.Atoi(strings.TrimSpace(rule[0]))
		if err != nil {
			fmt.Println(err)
		}
		right,err := strconv.Atoi(strings.TrimSpace(rule[1]))
		if err != nil {
			fmt.Println(err)
		}
		pagesBefore[left] = append(pagesBefore[left], right)
	}
	
	for k,v := range pagesBefore {
		fmt.Printf("key: %d, value: %v\n", k, v)
	}
	
	seqs := strings.Split(text[1], "\n")

	count := 0
	
	// for i:=0; i<len(seqs); i++ {
	// 	seq := strings.Split(seqs[i], ",")

	// 	for j:=1; j<len(seq); j++ {
	// 		num,err := strconv.Atoi(strings.TrimSpace(seq[j]))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 		for k:=0; k<j; k++ {
	// 			numback,err := strconv.Atoi(strings.TrimSpace(seq[k]))
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}

	// 			for l:=0; l<len(pagesBefore[num]); l++ {
	// 				if pagesBefore[num][l] == numback {
	// 					fmt.Println("FALSE")
	// 					count++
	// 					break
	// 				}
	// 			}
	// 		}
			
	// 	}
	// }
	for _, seqLine := range seqs {
		sequence := strings.Split(seqLine, ",")
		seqNums := make([]int, len(sequence))
		for i, numStr := range sequence {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				fmt.Println("Error parsing sequence number:", err)
				continue
			}
			seqNums[i] = num
		}

		fmt.Println(seqNums)

		outer:
		for j:=1; j<len(seqNums); j++ {
			currNum := seqNums[j]
			for k:=0; k<j; k++ {
				prevNum := seqNums[k]
				for _,l := range pagesBefore[currNum] {
					if l == prevNum {
						count++
						break outer
					}
				}
			}
		}

		
		
	}

	fmt.Println("count value: ", count)
	
}