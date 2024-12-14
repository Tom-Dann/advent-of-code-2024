package main

import (
	"fmt"
	"utils"
)

func positiveMod(x, m int) int {
	return (x%m + m) % m
}

const w, h = 101, 103

func safetyFactor(lines []string, secs int) int {
	quad := [4]int{}
	for _, line := range lines {
		var sx, sy, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &sx, &sy, &vx, &vy)
		x, y := positiveMod(sx+secs*vx, w), positiveMod(sy+secs*vy, h)
		if mx, my := (w-1)/2, (h-1)/2; x != mx && y != my {
			quad[x/(mx+1)+2*(y/(my+1))]++
		}
	}
	return quad[0] * quad[1] * quad[2] * quad[3]
}

func solve(lines []string) {
	min, safest := safetyFactor(lines, 100), 100
	fmt.Println("Part 1:", min)
	for i := 0; i < w*h; i++ {
		if score := safetyFactor(lines, i); score < min {
			min, safest = score, i
		}
	}
	fmt.Println("Part 2:", safest)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n"))
}
