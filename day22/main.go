package main

import (
	"fmt"
	"strconv"
	"utils"
)

func nextSecret(n int) int {
	n = (n ^ (n << 6)) % (1 << 24)
	n = (n ^ (n >> 5)) % (1 << 24)
	n = (n ^ (n << 11)) % (1 << 24)
	return n
}

func solve(lines []string) {
	part1, part2 := 0, 0
	bananas := map[[4]int]int{}
	for _, s := range lines {
		secret, _ := strconv.Atoi(s)
		seen, key := map[[4]int]bool{}, [4]int{0, 0, 0, 0}
		prev := secret % 10
		for i := 1; i <= 2000; i++ {
			secret = nextSecret(secret)
			curr := secret % 10
			key = [4]int{key[1], key[2], key[3], curr - prev}
			if i >= 4 && !seen[key] {
				bananas[key] += curr
				seen[key] = true
			}
			prev = curr
		}
		part1 += secret
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
