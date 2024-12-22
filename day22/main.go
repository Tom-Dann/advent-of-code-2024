package main

import (
	"fmt"
	"strconv"
	"utils"
)

func nextSecret(n int) int {
	n = (n ^ (n << 6)) % 16777216
	n = (n ^ (n >> 5)) % 16777216
	n = (n ^ (n << 11)) % 16777216
	return n
}

func solve(lines []string) {
	part1, part2 := 0, 0
	bananas := map[[4]int]int{}
	for _, s := range lines {
		secret, _ := strconv.Atoi(s)
		prices := [2001]int{secret % 10}
		for i := 1; i <= 2000; i++ {
			secret = nextSecret(secret)
			prices[i] = secret % 10
			part1 += secret
		}
		seen := map[[4]int]bool{}
		for i := 4; i < 2001; i++ {
			key := [4]int{
				prices[i-3] - prices[i-4],
				prices[i-2] - prices[i-3],
				prices[i-1] - prices[i-2],
				prices[i] - prices[i-1],
			}
			if !seen[key] {
				bananas[key] += prices[i]
				seen[key] = true
			}
		}
	}
	for _, count := range bananas {
		if count > part2 {
			part2 = count
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n"))
}
