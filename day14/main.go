package main

import (
	"fmt"
	"utils"
)

func positiveMod(x, m int) int {
	return (x%m + m) % m
}

const w, h = 101, 103

func quadCount(lines []string, secs int) [4]int {
	quad := [4]int{}
	for _, line := range lines {
		var sx, sy, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &sx, &sy, &vx, &vy)
		x, y := positiveMod(sx+secs*vx, w), positiveMod(sy+secs*vy, h)
		if mx, my := (w-1)/2, (h-1)/2; x != mx && y != my {
			quad[x/(mx+1)+2*(y/(my+1))]++
		}
	}
	return quad
}

func part1(lines []string) {
	quad := quadCount(lines, 100)
	fmt.Println("Part 1:", quad[0]*quad[1]*quad[2]*quad[3])
}

func part2(lines []string) {
	max, best := 0, 0
	for i := 0; i < w*h; i++ {
		for _, n := range quadCount(lines, i) {
			if n > max {
				max, best = n, i
			}
		}
	}
	fmt.Println("Part 2:", best)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, input)
	utils.TimeFunctionInput(part2, input)
}
