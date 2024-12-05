package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sortit(x [1000]int) ([1000]int){
	for i:=0; i<1000; i++{
		for j:=0; j<1000; j++{
			if x[i]<x[j] {
				x[i],x[j] = x[j],x[i]
			}
		}
	}
	return x
}

func sumD(x [1000]int, y [1000]int) (float64) {
	sum := 0.0
	
	for i:=0; i<1000; i++{
		sum += math.Abs(float64(x[i] - y[i]))
	}

	return sum
}



func main() {
	var left, right [1000]int

	
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}
	
	arroflines := strings.Split(string(data), "\n")
	
	for i:=0; i<1000; i++ {
		nums := strings.Split(arroflines[i], "   ")
		
		left[i], err = strconv.Atoi(nums[0])
		if err!=nil {
			fmt.Println(err)
			return
		}
		right[i], err = strconv.Atoi(strings.TrimSpace(nums[1]))
		if err!=nil {
			fmt.Println(err)
			return
		}
	}
	
	right = sortit(right)
	left = sortit(left)

	woahsum := sumD(left, right)

	fmt.Println("left array: ", left,"\nright array: ", right)
	fmt.Println("woahsum Value: ", woahsum)
}