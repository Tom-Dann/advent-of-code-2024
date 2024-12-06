package main

import (
	"fmt"
	"utils"
)

type Position struct {
	x, y int
}

func parseGrid(input []string) (map[Position]bool, Position) {
	grid := map[Position]bool{}
	pos := Position{}
	for j, line := range input {
		for i, char := range line {
			grid[Position{i, j}] = char == '#'
			if char == '^' {
				pos = Position{i, j}
			}
		}
	}
	return grid, pos
}

func move(grid map[Position]bool, added Position, pos Position, part2 chan<- int) map[Position]struct{} {
	type key struct{ p, d Position }
	dir := Position{0, -1}
	locs, seen := map[Position]struct{}{}, map[key]bool{}
	for {
		if part2 == nil {
			locs[pos] = struct{}{}
		} else {
			seen[key{pos, dir}] = true
		}
		new := Position{pos.x + dir.x, pos.y + dir.y}
		wall, valid := grid[new]
		if !valid {
			if part2 != nil {
				part2 <- 0
			}
			return locs
		}
		if wall || new == added {
			dir = Position{-dir.y, dir.x}
		} else {
			pos = new
		}
		if seen[key{pos, dir}] {
			part2 <- 1
			return locs
		}
	}
}

func solve(input []string) {
	grid, start := parseGrid(input)
	seen := move(grid, Position{-1, -1}, start, nil)
	fmt.Println("Part 1:", len(seen))

	count, part2 := 0, make(chan int, len(seen))
	for try := range seen {
		go move(grid, try, start, part2)
	}
	for i := 0; i < len(seen); i++ {
		count += <-part2
	}
	fmt.Println("Part 2:", count)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, input)
}
