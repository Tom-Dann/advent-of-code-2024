package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"utils"
)

func parseOrderings(input string) map[string]struct{} {
	lines := strings.Split(input, "\n")
	orderings := map[string]struct{}{}
	for _, line := range lines {
		orderings[line] = struct{}{}
	}
	return orderings
}

func parseLists(input string) [][]int {
	out := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		arr := []int{}
		for _, s := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(s)
			arr = append(arr, n)
		}
		out = append(out, arr)
	}
	return out
}

func getKey(a int, b int) string {
	return fmt.Sprintf("%d|%d", a, b)
}

func validOrder(orderings map[string]struct{}, list []int) bool {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			_, found := orderings[getKey(list[j], list[i])]
			if found {
				return false
			}
		}
	}
	return true
}

func compareFun(orderings map[string]struct{}) func(int, int) int {
	return func(a int, b int) int {
		_, less := orderings[getKey(a, b)]
		if less {
			return -1
		}
		_, more := orderings[getKey(b, a)]
		if more {
			return 1
		}
		return 0
	}
}

func solve(input []string) {
	orderings := parseOrderings(input[0])
	lists := parseLists(input[1])
	part1, part2 := 0, 0
	for _, list := range lists {
		if validOrder(orderings, list) {
			part1 += list[(len(list)-1)/2]
		} else {
			slices.SortFunc(list, compareFun(orderings))
			part2 += list[(len(list)-1)/2]
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	input := utils.ReadInput("input.txt", "\n\n")
	utils.TimeFunctionInput(solve, input)
}
