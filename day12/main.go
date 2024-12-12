package main

import (
	"fmt"
	"utils"
)

type Point struct {
	x, y int
}

func getNeighbours(p Point) []Point {
	return []Point{{p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}, {p.x, p.y - 1}}
}

func getCorner(p Point) [][3]Point {
	return [][3]Point{
		{{p.x + 1, p.y}, {p.x + 1, p.y + 1}, {p.x, p.y + 1}}, // ↘️
		{{p.x, p.y + 1}, {p.x - 1, p.y + 1}, {p.x - 1, p.y}}, // ↙️
		{{p.x - 1, p.y}, {p.x - 1, p.y - 1}, {p.x, p.y - 1}}, // ↖️
		{{p.x, p.y - 1}, {p.x + 1, p.y - 1}, {p.x + 1, p.y}}, // ↗️
	}
}

func solve(lines []string) {
	grid := map[Point]rune{}
	for j, line := range lines {
		for i, char := range line {
			grid[Point{i, j}] = char
		}
	}
	seen := map[Point]bool{}
	part1, part2 := 0, 0
	for pos, plant := range grid {
		if !seen[pos] {
			seen[pos] = true
			region, toVisit := map[Point]struct{}{pos: {}}, []Point{pos}
			corners, fences := 0, 0
			var curr Point
			for len(toVisit) > 0 {
				curr, toVisit = toVisit[0], toVisit[1:]
				for _, neighbour := range getNeighbours(curr) { // Part 1 - check direct neighbours and flood fill region
					if grid[neighbour] == plant {
						if !seen[neighbour] {
							region[neighbour], seen[neighbour] = struct{}{}, true
							toVisit = append(toVisit, neighbour)
						}
					} else {
						fences++
					}
				}
				for _, corner := range getCorner(curr) { // Part 2 - check for corners
					if grid[corner[0]] == plant && grid[corner[2]] == plant && grid[corner[1]] != plant { // 270° corner
						corners++
					} else if grid[corner[0]] != plant && grid[corner[2]] != plant { // 90° corner
						corners++
					}
				}
			}
			part1 += len(region) * fences
			part2 += len(region) * corners
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, input)
}
