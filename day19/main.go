package main

import (
	"fmt"
	"strings"
	"utils"
)

func solve(input []string) {
	possible, towels := map[string]bool{}, map[string]bool{}
	for _, towel := range strings.Split(input[0], ", ") {
		possible[towel] = true
		towels[towel] = true
	}
	var canMakePattern func(pattern string) bool
	canMakePattern = func(pattern string) bool {
		result, seen := possible[pattern]
		if seen {
			return result
		}
		for i := len(pattern) - 1; i > 0; i-- {
			if possible[pattern[i:]] && canMakePattern(pattern[:i]) {
				possible[pattern] = true
				return true
			}
		}
		possible[pattern] = false
		return false
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
		if canMakePattern(pattern) {
			part1++
			part2 += combos(pattern)
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n\n"))
}
