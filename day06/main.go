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

func doesItLoop(grid map[Position]bool, added Position, start Position, results chan<- int) {
	type key struct{ p, d Position }
	dir, pos := Position{0, -1}, start
	seen := map[key]struct{}{}
	for {
		seen[key{pos, dir}] = struct{}{}
		newPos := Position{pos.x + dir.x, pos.y + dir.y}
		wall, inGrid := grid[newPos]
		if !inGrid {
			results <- 0
			return
		}
		if wall || newPos == added {
			dir = Position{-dir.y, dir.x}
		} else {
			pos = newPos
		}
		_, loop := seen[key{pos, dir}]
		if loop {
			results <- 1
			return
		}
	}
}

func solve(input []string) {
	grid, start := parseGrid(input)
	pos, inGrid, dir := start, true, Position{0, -1}
	seen := map[Position]struct{}{}
	for inGrid {
		seen[pos] = struct{}{}
		newPos, wall := Position{pos.x + dir.x, pos.y + dir.y}, false
		wall, inGrid = grid[newPos]
		if wall {
			dir = Position{-dir.y, dir.x}
		} else {
			pos = newPos
		}
	}
	fmt.Println("Part 1:", len(seen))

	count, results := 0, make(chan int, len(seen))
	for try := range seen {
		if try != start {
			go doesItLoop(grid, try, start, results)
		} else {
			results <- 0
		}
	}
	for i := 0; i < len(seen); i++ {
		count += <-results
	}
	fmt.Println("Part 2:", count)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, input)
}
