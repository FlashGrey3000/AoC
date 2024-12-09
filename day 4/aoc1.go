package main

import (
	"fmt"
	"os"
	"strings"
)

func snowCrystal(i, j int, mat []string) (bool, int) {
	cnt := 0
	rows := len(mat)
	cols := len(mat[0])

	if j+3<cols && string(mat[i][j])+string(mat[i][j+1])+string(mat[i][j+2])+string(mat[i][j+3]) == "XMAS" {
		cnt++
	}

	if j-3>=0 && string(mat[i][j])+string(mat[i][j-1])+string(mat[i][j-2])+string(mat[i][j-3]) == "XMAS" {
		cnt++
	}

	if i+3<rows && string(mat[i][j])+string(mat[i+1][j])+string(mat[i+2][j])+string(mat[i+3][j]) == "XMAS" {
		cnt++
	}

	if i-3>=0 && string(mat[i][j])+string(mat[i-1][j])+string(mat[i-2][j])+string(mat[i-3][j]) == "XMAS" {
		cnt++
	}

	if i+3<rows && j+3<cols && string(mat[i][j])+string(mat[i+1][j+1])+string(mat[i+2][j+2])+string(mat[i+3][j+3]) == "XMAS" {
		cnt++
	}

	if i-3>=0 && j-3>=0 && string(mat[i][j])+string(mat[i-1][j-1])+string(mat[i-2][j-2])+string(mat[i-3][j-3]) == "XMAS" {
		cnt++
	}

	if i+3<rows && j-3>=0 && string(mat[i][j])+string(mat[i+1][j-1])+string(mat[i+2][j-2])+string(mat[i+3][j-3]) == "XMAS" {
		cnt++
	}

	if i-3>=0 && j+3<cols && string(mat[i][j])+string(mat[i-1][j+1])+string(mat[i-2][j+2])+string(mat[i-3][j+3]) == "XMAS" {
		cnt++
	}

	return cnt > 0, cnt
}


func main() {
	data, err := os.ReadFile("input.txt")
	if err!=nil {
		fmt.Println(err)
	}
	mat := strings.Split(string(data), "\n")
	count := 0
	for i:=0; i<len(mat); i++ {
		for j:=0; j<len(mat[0]); j++ {
			if string(mat[i][j]) != "X" {
				continue
			}
			hasMas, cnt := snowCrystal(i, j, mat) 
			if hasMas {
				count += cnt
			}
		}
		fmt.Printf("loop %d done...\n", i)
	}

	fmt.Println("Occurence of XMAS: ", count)
	
}