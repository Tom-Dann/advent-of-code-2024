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

func part1(input []string) {
	grid, pos := parseGrid(input)
	inGrid, dir := true, Position{0, -1}
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
}

func doesItLoop(grid map[Position]bool, added Position, start Position) bool {
	type key struct{ p, d Position }
	dir, pos := Position{0, -1}, start
	seen := map[key]struct{}{}
	for {
		seen[key{pos, dir}] = struct{}{}
		newPos := Position{pos.x + dir.x, pos.y + dir.y}
		wall, valid := grid[newPos]
		if !valid {
			return false
		}
		if wall || (newPos.x == added.x && newPos.y == added.y) {
			dir = Position{-dir.y, dir.x}
		} else {
			pos = newPos
		}
		_, loop := seen[key{pos, dir}]
		if loop {
			return true
		}
	}
}

func part2(input []string) {
	grid, pos := parseGrid(input)
	count := 0
	for try, wall := range grid {
		if !(wall || (try.x == pos.x && try.y == pos.y)) {
			if doesItLoop(grid, try, pos) {
				count++
			}
		}
	}
	fmt.Println("Part 2:", count)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, input)
	utils.TimeFunctionInput(part2, input)
}
