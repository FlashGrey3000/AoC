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
	//seqs := strings.Split(text[1], "\n")

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

	
	
}