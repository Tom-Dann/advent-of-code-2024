package main

import (
	"fmt"
	"utils"
)

type Position struct{ x, y int }

func (p Position) inBounds(h, w int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < w && p.y < h
}

func nodes(a, b Position) []Position {
	return []Position{{2*a.x - b.x, 2*a.y - b.y}, {2*b.x - a.x, 2*b.y - a.y}}
}

func nodesColinear(a, b Position, h, w int) []Position {
	dx, dy := a.x-b.x, a.y-b.y
	curr := a
	nodes := []Position{}
	for curr.inBounds(h, w) {
		nodes = append(nodes, curr)
		curr = Position{curr.x + dx, curr.y + dy}
	}
	curr = b
	for curr.inBounds(h, w) {
		nodes = append(nodes, curr)
		curr = Position{curr.x - dx, curr.y - dy}
	}
	return nodes
}

func solve(lines []string) {
	h, w := len(lines), len(lines[0])
	antennas := map[rune][]Position{}
	part1, part2 := map[Position]struct{}{}, map[Position]struct{}{}
	for j, line := range lines {
		for i, char := range line {
			if char != '.' {
				arr, compare := antennas[char]
				curr := Position{i, j}
				if compare {
					for _, antenna := range arr {
						for _, node := range nodes(curr, antenna) {
							if node.inBounds(h, w) {
								part1[node] = struct{}{}
							}
						}
						for _, node := range nodesColinear(curr, antenna, h, w) {
							part2[node] = struct{}{}
						}
					}
					antennas[char] = append(antennas[char], curr)
				} else {
					antennas[char] = []Position{curr}
				}
			}
		}
	}
	fmt.Println("Part 1:", len(part1))
	fmt.Println("Part 2:", len(part2))
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, input)
}
