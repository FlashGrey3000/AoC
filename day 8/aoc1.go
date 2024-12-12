package main

import (
	"fmt"
	// "os"

	utils "github.com/FlashGrey3000/AoC/utils"
)

type point struct{
	X, Y int
}

func main() {
	lines := utils.ReadLines("input.txt")

	grid := make(map[byte][]point, 0)

	for y,line := range lines {
		for x,b := range line {
			grid[byte(b)] = append(grid[byte(b)], point{X: x, Y: y})
		}
	}

	fmt.Println(grid)

	fmt.Println("antinodes: ", calcAntiNodes(grid))
}

func calcAntiNodes(grid map[byte][]point) (int) {
	count := 0
	for k,v := range grid {
		if k == '.' {
			continue
		}
		for i:=0; i<len(v); i++ {
			for j:=0; j<len(v); j++ {
				if i==j {
					continue
				}
				x:=v[j].X + (v[j].X - v[i].X)
				y:=v[j].Y + (v[j].Y - v[i].Y)
				fmt.Print(v[j].X, v[j].Y, v[i].X, v[i].Y)
				if x >= 0 && x < 12 && y >= 0 && y < 12 {
					count++
					fmt.Print(" : ", x, y)
				}
				fmt.Print("\n")
			}
		}
	}

	return count
}