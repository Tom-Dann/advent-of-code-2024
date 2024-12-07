package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"utils"
)

func parseLine(line string) (int64, []int64) {
	nums := []int64{}
	split := strings.Split(line, ": ")
	total, _ := strconv.ParseInt(split[0], 10, 64)
	for _, s := range strings.Split(split[1], " ") {
		n, _ := strconv.ParseInt(s, 10, 64)
		nums = append(nums, n)
	}
	return total, nums
}

func valid(total int64, nums []int64, ops int) bool {
	for i := 0; i < int(math.Pow(float64(ops), float64(len(nums)-1))); i++ {
		calc := nums[0]
		for j := 0; j < len(nums)-1; j++ {
			switch (i / int(math.Pow(float64(ops), float64(j)))) % ops {
			case 0:
				calc += nums[j+1]
			case 1:
				calc *= nums[j+1]
			case 2:
				s := strconv.FormatInt(calc, 10) + strconv.FormatInt(nums[j+1], 10)
				calc, _ = strconv.ParseInt(s, 10, 64)
			}
		}
		if calc == total {
			return true
		}
	}
	return false
}

func solve(lines []string) {
	part1, part2 := int64(0), int64(0)
	for _, line := range lines {
		total, nums := parseLine(line)
		if valid(total, nums, 2) {
			part1 += total
			part2 += total
		} else if valid(total, nums, 3) {
			part2 += total
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, input)
}
