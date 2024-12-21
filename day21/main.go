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
	moves := ""
	dx, dy := b.x-a.x, b.y-a.y
	if a.x == avoid.x && b.y == avoid.y {
		if dx < 0 {
			moves += strings.Repeat("<", -dx)
		}
		if dx > 0 {
			moves += strings.Repeat(">", dx)
		}
		if dy < 0 {
			moves += strings.Repeat("^", -dy)
		}
		if dy > 0 {
			moves += strings.Repeat("v", dy)
		}
	} else if a.y == avoid.y && b.x == avoid.x {
		if dy < 0 {
			moves += strings.Repeat("^", -dy)
		}
		if dy > 0 {
			moves += strings.Repeat("v", dy)
		}
		if dx < 0 {
			moves += strings.Repeat("<", -dx)
		}
		if dx > 0 {
			moves += strings.Repeat(">", dx)
		}
	} else {
		if dx < 0 {
			moves += strings.Repeat("<", -dx)
		}
		if dy < 0 {
			moves += strings.Repeat("^", -dy)
		}
		if dy > 0 {
			moves += strings.Repeat("v", dy)
		}
		if dx > 0 {
			moves += strings.Repeat(">", dx)
		}
	}
	return moves + "A"
}

type State struct {
	seq   string
	level int
}

var cache = map[State]int{}

func getLength(seq string, level int) int {
	key := State{seq, level}
	result, seen := cache[key]
	if seen {
		return result
	}
	if level == 0 {
		return len(seq)
	}
	length := 0
	curr := 'A'
	for _, next := range seq {
		length += getLength(getMoves(robotPad[curr], robotPad[next], Point{0, 0}), level-1)
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
		curr := keyPad['A']
		sequence := ""
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
