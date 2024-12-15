package main

import (
	"fmt"
	"strings"
	"utils"
)

type Point struct {
	x, y int
}

var direction = map[rune]Point{
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
	'^': {0, -1},
}

func gpsScore(grid map[Point]rune, box rune) int {
	score := 0
	for pos, char := range grid {
		if char == box {
			score += pos.x + 100*(pos.y)
		}
	}
	return score
}

func part1(input []string) {
	grid := map[Point]rune{}
	var robot Point
	for j, line := range strings.Split(input[0], "\n") {
		for i, char := range line {
			grid[Point{i, j}] = char
			if char == '@' {
				robot = Point{i, j}
			}
		}
	}
	for _, instr := range strings.ReplaceAll(input[1], "\n", "") {
		dir := direction[instr]
		next := Point{robot.x + dir.x, robot.y + dir.y}
		check := next
		for grid[check] == 'O' {
			check.x += dir.x
			check.y += dir.y
		}
		if grid[check] == '.' {
			grid[robot], grid[check], grid[next] = '.', 'O', '@'
			robot = next
		}
	}
	fmt.Println("Part 1:", gpsScore(grid, 'O'))
}

func part2(input []string) {
	grid := map[Point]rune{}
	var robot Point
	for j, line := range strings.Split(input[0], "\n") {
		for i, char := range line {
			switch char {
			case '.':
				fallthrough
			case '#':
				grid[Point{2 * i, j}] = char
				grid[Point{2*i + 1, j}] = char
			case 'O':
				grid[Point{2 * i, j}] = '['
				grid[Point{2*i + 1, j}] = ']'
			case '@':
				grid[Point{2 * i, j}] = '@'
				grid[Point{2*i + 1, j}] = '.'
				robot = Point{2 * i, j}
			}
		}
	}
	for _, instr := range strings.ReplaceAll(input[1], "\n", "") {
		dir := direction[instr]
		if instr == '<' || instr == '>' {
			check := Point{robot.x + dir.x, robot.y + dir.y}
			for grid[check] == '[' || grid[check] == ']' {
				check.x += dir.x
				check.y += dir.y
			}
			if grid[check] == '.' {
				for check != robot {
					next := Point{check.x - dir.x, check.y - dir.y}
					grid[check] = grid[next]
					check = next
				}
				grid[robot] = '.'
				robot = Point{robot.x + dir.x, robot.y + dir.y}
			}
		} else {
			pushing, wall := true, false
			checks := map[Point]struct{}{robot: {}}
			moves := map[Point]rune{}
			for pushing && !wall {
				nextChecks := map[Point]struct{}{}
				pushing = false
				for check := range checks {
					next := Point{check.x + dir.x, check.y + dir.y}
					if grid[next] == '#' {
						wall = true
						break
					}
					if grid[next] == ']' || grid[next] == '[' {
						pushing = true
						nextChecks[next] = struct{}{}
						if grid[next] == ']' {
							nextChecks[Point{next.x - 1, next.y}] = struct{}{}
						} else {
							nextChecks[Point{next.x + 1, next.y}] = struct{}{}
						}
					}
					moves[next] = grid[check]
					_, found := moves[check]
					if !found {
						moves[check] = '.'
					}
				}
				checks = nextChecks
			}
			if !wall {
				for pos, new := range moves {
					grid[pos] = new
				}
				robot = Point{robot.x + dir.x, robot.y + dir.y}
			}
		}
	}
	fmt.Println("Part 2:", gpsScore(grid, '['))
}

func main() {
	input := utils.ReadInput("input.txt", "\n\n")
	utils.TimeFunctionInput(part1, input)
	utils.TimeFunctionInput(part2, input)
}
