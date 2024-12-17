package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func runProgram(vals [7]int, ops []int) string {
	index, output := 0, ""
	for index < len(ops) {
		literal := ops[index+1]
		combo := vals[literal]
		switch ops[index] {
		case 0: // adv
			vals[4] = vals[4] >> combo
		case 1: // bxl
			vals[5] = vals[5] ^ literal
		case 2: // bst
			vals[5] = combo % 8
		case 3: //jnz
			if vals[4] != 0 {
				index = literal - 2
			}
		case 4: // bxc
			vals[5] = vals[5] ^ vals[6]
		case 5: // out
			output += "," + strconv.Itoa(combo%8)
		case 6: // bdv
			vals[5] = vals[4] >> combo
		case 7: // cdv
			vals[6] = vals[4] >> combo
		}
		index += 2
	}
	return output[1:]
}

func octet(a []int) int {
	out, mul := 0, 1
	for i := range a {
		out += a[len(a)-i-1] * mul
		mul *= 8
	}
	return out
}

func solve(lines []string) {
	vals := [7]int{0, 1, 2, 3}
	var opsString string
	fmt.Sscanf(lines[0], "Register A: %d\nRegister B: %d\nRegister C: %d\n\nProgram: %s", &vals[4], &vals[5], &vals[6], &opsString)
	ops := []int{}
	for _, s := range strings.Split(opsString, ",") {
		n, _ := strconv.Atoi(s)
		ops = append(ops, n)
	}
	fmt.Println("Part 1:", runProgram(vals, ops))

	curr := []int{}
	for j := 0; j < len(ops); j++ {
		for i := 0; true; i++ {
			code := append(curr, i)
			vals[4] = octet(code)
			if strings.HasSuffix(opsString, runProgram(vals, ops)) {
				curr = code
				break
			}
		}
	}
	fmt.Println("Part 2:", octet(curr))
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "🙂"))
}
