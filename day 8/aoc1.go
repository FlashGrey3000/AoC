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

	fmt.Println("antinodes: ", calcAntiNodes(grid))
	fmt.Println("Antinodes after considering resonance: ", ResonantAntiNodes(grid))
}

func calcAntiNodes(grid map[byte][]point) (int) {
	count := 0
	visited := make([]point, 0)
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
				if x >= 0 && x < 50 && y >= 0 && y < 50 && !overlapping(x,y, visited){
					count++
					visited = append(visited, point{X: x, Y: y})
				}
			}
		}
	}

	return count
}

func overlapping(x, y int, visited []point) (bool) {
	p := point{X: x, Y: y}
	for _,po := range visited {
		if po == p {
			return true
		}
	}
	return false
}

func ResonantAntiNodes(grid map[byte][]point) (int) {
	count := 0
	visited := make([]point, 0)
	for k,v := range grid {
		if k == '.' {
			continue
		}
		for i:=0; i<len(v); i++ {
			for j:=0; j<len(v); j++ {
				if i==j {
					continue
				}
				for z:=1; ; z++ {
					x:=v[j].X + z*(v[j].X - v[i].X)
					y:=v[j].Y + z*(v[j].Y - v[i].Y)
					if x >= 0 && x < 50 && y >= 0 && y < 50 && !overlapping(x,y, visited){
						count++
						visited = append(visited, point{X: x, Y: y})
					} else {
						break
					}
				}
			}
		}
	}

	return count
}