package main

import (
	"fmt"
	"strconv"
	"utils"
)

func parseInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func solve(input []string) {
	counts, totals := map[int64]int64{}, [2]int64{}
	for _, s := range input {
		counts[parseInt(s)]++
	}
	for i := 1; i <= 75; i++ {
		newCount := map[int64]int64{}
		for k, v := range counts {
			if k == 0 {
				newCount[1] += v
			} else if s := strconv.FormatInt(k, 10); len(s)%2 == 0 {
				newCount[parseInt(s[:len(s)/2])] += v
				newCount[parseInt(s[len(s)/2:])] += v
			} else {
				newCount[k*2024] += v
			}
		}
		counts = newCount
		if i == 25 || i == 75 {
			for _, v := range counts {
				totals[i/75] += v
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
