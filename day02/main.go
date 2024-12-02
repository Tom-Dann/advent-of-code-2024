package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func parseLine(line string) []int {
	vals := strings.Fields(line)
	arr := make([]int, len(vals))
	for i, s := range vals {
		arr[i], _ = strconv.Atoi(s)
	}
	return arr
}

func removeIndex(arr []int, i int) []int {
	new := append([]int{}, arr[:i]...)
	return append(new, arr[i+1:]...)
}

func safe(nums []int) bool {
	increasing := nums[0] < nums[1]
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if d > 0 != increasing || d < -3 || d == 0 || d > 3 {
			return false
		}
	}
	return true
}

func safeWithTolerance(nums []int) bool {
	increasing := nums[0] < nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if d > 0 != increasing || d < -3 || d == 0 || d > 3 {
			return safe(removeIndex(nums, i-1)) || safe(removeIndex(nums, i))
		}
	}
	return true
}

func solve(lines []string) {
	part1, part2 := 0, 0
	for _, line := range lines {
		nums := parseLine(line)
		if safe(nums) {
			part1++
			part2++
		} else if safeWithTolerance(nums) {
			part2++
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, input)
}
