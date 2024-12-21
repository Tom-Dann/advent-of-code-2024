package main

import (
	"fmt"
	"strings"
	"utils"
)

type Point struct{ x, y int }

var keyPad = map[rune]Point{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'0': {1, 3},
	'A': {2, 3},
}

var robotPad = map[rune]Point{
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

func getMoves(a, b, avoid Point) string {
	var moves, xMove, yMove string
	dx, dy := b.x-a.x, b.y-a.y
	if dx >= 0 {
		xMove = strings.Repeat(">", dx)
	} else {
		xMove = strings.Repeat("<", -dx)
	}
	if dy >= 0 {
		yMove = strings.Repeat("v", dy)
	} else {
		yMove = strings.Repeat("^", -dy)
	}
	if (a.x == avoid.x && b.y == avoid.y) || (dx < 0 && !(a.y == avoid.y && b.x == avoid.x)) {
		moves = xMove + yMove
	} else {
		moves = yMove + xMove
	}
	return moves + "A"
}

type State struct {
	sequence string
	level    int
}

var cache = map[State]int{}

func getLength(sequence string, level int) int {
	key := State{sequence, level}
	if result, seen := cache[key]; seen {
		return result
	}
	if level == 0 {
		return len(sequence)
	}
	length, curr := 0, robotPad['A']
	for _, button := range sequence {
		next := robotPad[button]
		length += getLength(getMoves(curr, next, Point{0, 0}), level-1)
		curr = next
	}
	cache[key] = length
	return length
}

func solve(lines []string) {
	part1, part2 := 0, 0
	for _, code := range lines {
		var codeNum int
		fmt.Sscanf(code, "%dA", &codeNum)
		sequence, curr := "", keyPad['A']
		for _, button := range code {
			next := keyPad[button]
			sequence += getMoves(curr, next, Point{0, 3})
			curr = next
		}
		part1 += codeNum * getLength(sequence, 2)
		part2 += codeNum * getLength(sequence, 25)
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n"))
}
