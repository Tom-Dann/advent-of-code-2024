package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"utils"
)

func parseLine(line string) (int, int) {
	nums := strings.Fields(line)
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	return x, y
}

func parseLists(lines []string) ([]int, []int) {
	size := len(lines)
	a := make([]int, size)
	b := make([]int, size)
	for i, line := range lines {
		a[i], b[i] = parseLine(line)
	}
	return a, b
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func part1(lines []string) {
	a, b := parseLists(lines)
	slices.Sort(a)
	slices.Sort(b)

	sum := 0
	for i := range a {
		sum += abs(a[i] - b[i])
	}
	fmt.Println("Part 1:", sum)
}

func part2(lines []string) {
	a, b := parseLists(lines)
	counts := make(map[int]int)
	for _, n := range b {
		counts[n]++
	}
	sum := 0
	for _, n := range a {
		sum += counts[n] * n
	}
	fmt.Println("Part 2:", sum)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, input)
	utils.TimeFunctionInput(part2, input)
}
