package main

import (
	"fmt"
	"strconv"
	"utils"
)

func solve(input []string) {
	counts := map[int64]int64{}
	totals := [2]int64{}
	for _, s := range input {
		n, _ := strconv.ParseInt(s, 10, 64)
		counts[n]++
	}
	for i := 1; i <= 75; i++ {
		newCount := map[int64]int64{}
		for k, v := range counts {
			if k == 0 {
				newCount[1] += v
			} else if s := strconv.FormatInt(k, 10); len(s)%2 == 0 {
				a, _ := strconv.ParseInt(s[:len(s)/2], 10, 64)
				b, _ := strconv.ParseInt(s[len(s)/2:], 10, 64)
				newCount[a] += v
				newCount[b] += v
			} else {
				newCount[k*2024] += v
			}
		}
		counts = newCount
		if i == 25 || i == 75 {
			for _, v := range counts {
				totals[(i-25)/50] += v
			}
		}
	}
	fmt.Println("Part 1:", totals[0])
	fmt.Println("Part 2:", totals[1])
}

func main() {
	input := utils.ReadInput("input.txt", " ")
	utils.TimeFunctionInput(solve, input)
}
