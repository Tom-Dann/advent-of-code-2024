package main

import (
	"fmt"
	"utils"
)

func getCost(Ax, Ay, Bx, By, Px, Py int) int {
	n, d := By*Px-Bx*Py, By*Ax-Bx*Ay
	if n%d == 0 {
		i := n / d
		return 3*i + ((Py - i*Ay) / By)
	}
	return 0
}

func solve(input []string) {
	part1, part2 := 0, 0
	for _, section := range input {
		var Ax, Ay, Bx, By, Px, Py int
		fmt.Sscanf(section, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d", &Ax, &Ay, &Bx, &By, &Px, &Py)
		part1 += getCost(Ax, Ay, Bx, By, Px, Py)
		part2 += getCost(Ax, Ay, Bx, By, Px+1e13, Py+1e13)
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	input := utils.ReadInput("input.txt", "\n\n")
	utils.TimeFunctionInput(solve, input)
}
