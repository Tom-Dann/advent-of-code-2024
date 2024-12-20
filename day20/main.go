package main

import (
	"fmt"
	"utils"
)

type Point struct{ x, y int }

func getNeighbours(p Point) []Point {
	return []Point{{p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}, {p.x, p.y - 1}}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func solve(lines []string) {
	grid := map[Point]rune{}
	var curr Point
	for j, line := range lines { // Parse grid
		for i, char := range line {
			grid[Point{i, j}] = char
			if char == 'S' {
				curr = Point{i, j}
			}
		}
	}

	positions := map[Point]int{curr: 0}
	path := map[Point]bool{curr: true}
	for grid[curr] != 'E' { // Find path from start (S) to end (E)
		for _, next := range getNeighbours(curr) {
			if grid[next] != '#' && !path[next] {
				path[next] = true
				positions[next] = positions[curr] + 1
				curr = next
				break
			}
		}
	}

	part1, part2 := 0, 0
	for startCheat, startIndex := range positions { // Find all cheats from the path to another path point within manhattan dist 20
		for i := startCheat.x - 20; i <= startCheat.x+20; i++ {
			xDist := abs(startCheat.x - i)
			for j := startCheat.y - (20 - xDist); j <= startCheat.y+(20-xDist); j++ {
				yDist := abs(startCheat.y - j)
				if endIndex := positions[Point{i, j}]; endIndex >= startIndex+100+xDist+yDist {
					if xDist+yDist <= 2 { // Manhattan dist <= 2 for part 1
						part1++
					}
					part2++
				}
			}
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n"))
}
