package main

import (
	"fmt"
	"utils"
)

type Position struct {
	x, y int
}

func parseGrid(lines []string) map[Position]int {
	grid := map[Position]int{}
	for j, line := range lines {
		for i := range line {
			grid[Position{i, j}] = int(line[i]) - 48
		}
	}
	return grid
}

func getNeighbours(p Position) []Position {
	return []Position{{p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}, {p.x, p.y - 1}}
}

func part1(lines []string) {
	grid, total := parseGrid(lines), 0
	for pos := range grid {
		if grid[pos] == 0 {
			toCheck := map[Position]struct{}{pos: {}}
			for height := 1; height <= 9; height++ {
				next := map[Position]struct{}{}
				for check := range toCheck {
					for _, neighbour := range getNeighbours(check) {
						if grid[neighbour] == height {
							next[neighbour] = struct{}{}
						}
					}
				}
				toCheck = next
			}
			total += len(toCheck)
		}
	}
	fmt.Println("Part 1:", total)
}

func hike(pos Position, grid map[Position]int) int {
	if grid[pos] == 9 {
		return 1
	}
	total := 0
	for _, neighbour := range getNeighbours(pos) {
		if grid[neighbour] == grid[pos]+1 {
			total += hike(neighbour, grid)
		}
	}
	return total
}

func part2(lines []string) {
	grid, total := parseGrid(lines), 0
	for pos := range grid {
		if grid[pos] == 0 {
			total += hike(pos, grid)
		}
	}
	fmt.Println("Part 2:", total)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, input)
	utils.TimeFunctionInput(part2, input)
}
