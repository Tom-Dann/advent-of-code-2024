package main

import (
	"fmt"
	"utils"
)

func schematicNumber(section string) int {
	num, mult := 0, 1
	for _, char := range section {
		if char == '#' {
			num += mult
		}
		mult <<= 1
	}
	return num
}

func solve(input []string) {
	var keys, locks []int
	for _, section := range input {
		value := schematicNumber(section)
		if section[0] == '.' {
			keys = append(keys, value)
		} else {
			locks = append(locks, value)
		}
	}
	total := 0
	for _, key := range keys {
		for _, lock := range locks {
			if key&lock == 0 {
				total++
			}
		}
	}
	fmt.Println("Part 1:", total)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n\n"))
}
