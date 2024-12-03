package main

import (
	"fmt"
	"regexp"
	"strconv"
	"utils"
)

func multMatch(match []string) int {
	a, _ := strconv.Atoi(match[1])
	b, _ := strconv.Atoi(match[2])
	return a * b
}

func solve(lines []string) {
	part1, part2, enabled := 0, 0, true
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do(?:n't)?\(\)`)
	matches := re.FindAllStringSubmatch(lines[0], -1)
	for _, match := range matches {
		switch match[0][2] {
		case '(': // do()
			enabled = true
		case 'n': // don't()
			enabled = false
		case 'l': // mul(a,b)
			x := multMatch(match)
			part1 += x
			if enabled {
				part2 += x
			}
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	input := utils.ReadInput("input.txt", "\n\n")
	utils.TimeFunctionInput(solve, input)
}
