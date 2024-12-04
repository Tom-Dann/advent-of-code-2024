package main

import (
	"fmt"
	"regexp"
	"utils"
)

func parseGrid(lines []string) [][]rune {
	grid := [][]rune{}
	for _, line := range lines {
		gridLine := []rune{}
		for _, r := range line {
			gridLine = append(gridLine, r)
		}
		grid = append(grid, gridLine)
	}
	return grid
}

func search(r []rune) int {
	reFor := regexp.MustCompile("XMAS")
	reRev := regexp.MustCompile("SAMX")
	s := string(r)
	return len(reFor.FindAllString(s, -1)) + len(reRev.FindAllString(s, -1))
}

func part1(lines []string) {
	grid := parseGrid(lines)
	sum, h, w := 0, len(grid), len(grid[0])
	for j := 0; j < h; j++ { // —
		sum += search(grid[j])
	}
	for i := 0; i < w; i++ { // |
		line := []rune{}
		for j := 0; j < h; j++ {
			line = append(line, grid[j][i])
		}
		sum += search(line)
	}
	for js := 0; js < h; js++ {
		diag, onal := []rune{}, []rune{}
		for i, j := 0, js; j < h && i < w; i, j = i+1, j+1 {
			diag = append(diag, grid[j][i])     // ◤
			onal = append(onal, grid[h-j-1][i]) // ◣
		}
		sum += search(diag) + search(onal)
	}
	for is := 1; is < w; is++ {
		diag, onal := []rune{}, []rune{}
		for i, j := is, 0; j < h && i < w; i, j = i+1, j+1 {
			diag = append(diag, grid[j][i])     // ◢
			onal = append(onal, grid[h-j-1][i]) // ◥
		}
		sum += search(diag) + search(onal)
	}
	fmt.Println("Part 1:", sum)
}

func checkCorners(a rune, b rune) bool {
	return (a == 'M' && b == 'S') || (a == 'S' && b == 'M')
}

func part2(lines []string) {
	grid := parseGrid(lines)
	sum, h, w := 0, len(grid), len(grid[0])
	for j := 1; j < h-1; j++ {
		for i := 1; i < w-1; i++ {
			if grid[j][i] == 'A' {
				a := checkCorners(grid[j-1][i-1], grid[j+1][i+1])
				b := checkCorners(grid[j+1][i-1], grid[j-1][i+1])
				if a && b {
					sum++
				}
			}
		}
	}
	fmt.Println("Part 2:", sum)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, input)
	utils.TimeFunctionInput(part2, input)
}
