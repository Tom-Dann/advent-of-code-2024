package main

import (
	"fmt"
	"strings"
	"utils"
)

func solve(input []string) {
	towels := map[string]bool{}
	for _, towel := range strings.Split(input[0], ", ") {
		towels[towel] = true
	}

	cache := map[string]int{}
	var combos func(pattern string) int
	combos = func(pattern string) int {
		result, seen := cache[pattern]
		if seen {
			return result
		}
		count := 0
		if towels[pattern] {
			count++
		}
		for i := 0; i < len(pattern); i++ {
			if towels[pattern[:i]] {
				count += combos(pattern[i:])
			}
		}
		cache[pattern] = count
		return count
	}

	part1, part2 := 0, 0
	for _, pattern := range strings.Split(input[1], "\n") {
		if count := combos(pattern); count > 0 {
			part1++
			part2 += count
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n\n"))
}
