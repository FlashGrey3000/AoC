package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data,err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	visited := make([][2]int, 0)

	location := strings.Split(string(data), "\r\n")

	loc := make([][]byte, len(location))
	og_loc := make([][]byte, len(location))
	for i,r := range location {
		loc[i] = []byte(r)
		og_loc[i] = []byte(r)
	}

	guard_x, guard_y := 59,77
	fmt.Println(string(loc[guard_y][guard_x]))
	visited = append(visited, [2]int{guard_x, guard_y})

	for {
		if loc[guard_y][guard_x] == '^' {
			if guard_y==0 {
				break
			}
			if loc[guard_y-1][guard_x] != '#' {
				loc[guard_y][guard_x] = '.'
				guard_y--
				loc[guard_y][guard_x] = '^'
			} else {
				loc[guard_y][guard_x] = '>'
			}
			isVisited := false
			for _, coord := range visited {
				if coord[0] == guard_x && coord[1] == guard_y {
					isVisited = true
					break
				}
			}
			if !isVisited {
				visited = append(visited, [2]int{guard_x, guard_y})
			}
		} else if loc[guard_y][guard_x] == '>' {
			if guard_x==129 {
				break
			}
			if loc[guard_y][guard_x+1] != '#' {
				loc[guard_y][guard_x] = '.'
				guard_x++
				loc[guard_y][guard_x] = '>'
			} else {
				loc[guard_y][guard_x] = 'v'
			}
			isVisited := false
			for _, coord := range visited {
				if coord[0] == guard_x && coord[1] == guard_y {
					isVisited = true
					break
				}
			}
			if !isVisited {
				visited = append(visited, [2]int{guard_x, guard_y})
			}
		} else if loc[guard_y][guard_x] == 'v' {
			if guard_y == 129 {
				break
			}
			if loc[guard_y+1][guard_x] != '#' {
				loc[guard_y][guard_x] = '.'
				guard_y++
				loc[guard_y][guard_x] = 'v'
			} else {
				loc[guard_y][guard_x] = '<'
			}
			isVisited := false
			for _, coord := range visited {
				if coord[0] == guard_x && coord[1] == guard_y {
					isVisited = true
					break
				}
			}
			if !isVisited {
				visited = append(visited, [2]int{guard_x, guard_y})
			}
		} else if loc[guard_y][guard_x] == '<' {
			if guard_x == 0 {
				break
			}
			if loc[guard_y][guard_x-1] != '#' {
				loc[guard_y][guard_x] = '.'
				guard_x--
				loc[guard_y][guard_x] = '<'
			} else {
				loc[guard_y][guard_x] = '^'
				
			}
			isVisited := false
			for _, coord := range visited {
				if coord[0] == guard_x && coord[1] == guard_y {
					isVisited = true
					break
				}
			}
			if !isVisited {
				visited = append(visited, [2]int{guard_x, guard_y})
			}
		} else {
			fmt.Println(loc[guard_y][guard_x])
		}
	}

	fmt.Println(len(visited))

	fmt.Println(string(og_loc[77][59]))
	count := part2(og_loc, visited)


	fmt.Println("part 2: ", count)
}

func part2(loc [][]byte, visited [][2]int) (int) {
	count := 0

	for i:=0; i<len(visited); i++ {
		if i==0 {
			continue
		}
		og_loc := make([][]byte, len(loc))
		for i,r := range loc {
			og_loc[i] = []byte(r)
		}
		if addObstacleAndCheckForLoop(loc, visited[i]) == true {
			count++
		}
		loc = og_loc
	}

	return count
}

func addObstacleAndCheckForLoop(loc [][]byte, pos [2]int) bool {
    original := loc[pos[1]][pos[0]]

    loc[pos[1]][pos[0]] = '#'
	
    guard_x, guard_y := 59, 77
    direction := loc[guard_y][guard_x]
	fmt.Println(string(direction))

    validDirections := map[byte]bool{'^': true, '>': true, 'v': true, '<': true}
    if !validDirections[direction] {
        fmt.Printf("Invalid initial direction at (%d, %d): %c\n", guard_x, guard_y, direction)
        loc[pos[1]][pos[0]] = original
        return false
    }

    visited := make(map[[3]int]struct{})

    for {
        fmt.Printf("Before move: x=%d, y=%d, dir=%c\n", guard_x, guard_y, direction)
        state := [3]int{guard_x, guard_y, int(direction)}
        if _, exists := visited[state]; exists {
            fmt.Println("Loop detected!")
            loc[pos[1]][pos[0]] = original
            return true
        }
        visited[state] = struct{}{}
        loc[guard_y][guard_x] = '.'

        switch direction {
        case '^':
            if guard_y > 0 && loc[guard_y-1][guard_x] != '#' {
                guard_y--
            } else {
                direction = '>'
            }
        case '>':
            if guard_x < len(loc[0])-1 && loc[guard_y][guard_x+1] != '#' {
                guard_x++
            } else {
                direction = 'v'
            }
        case 'v':
            if guard_y < len(loc)-1 && loc[guard_y+1][guard_x] != '#' {
                guard_y++
            } else {
                direction = '<'
            }
        case '<':
            if guard_x > 0 && loc[guard_y][guard_x-1] != '#' {
                guard_x--
            } else {
                direction = '^'
            }
        default:
            fmt.Printf("Invalid direction encountered at (%d, %d): %c\n", guard_x, guard_y, direction)
            loc[pos[1]][pos[0]] = original
            return false
        }
        loc[guard_y][guard_x] = byte(direction)

        if guard_x < 0 || guard_y < 0 || guard_x >= len(loc[0]) || guard_y >= len(loc) {
            break
        }
    }

    loc[pos[1]][pos[0]] = original
    return false
} // I am alive ;-;
