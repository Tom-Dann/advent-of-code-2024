package main

import (
	"fmt"
	"utils"
)

func toInt(b byte) int {
	return int(b) - 48
}

func part1() {
	input := utils.ReadFile("input.txt")
	total, start := int64(0), true
	pos, i, j := 0, 0, len(input)-1
	a, b := toInt(input[i]), toInt(input[j])
	for i < j || b > 0 {
		if a == 0 {
			i++
			start = !start
			a = toInt(input[i])
		} else if start {
			total += int64(pos * i / 2)
			a--
			pos++
		} else if b > 0 {
			total += int64(pos * j / 2)
			b--
			a--
			pos++
		} else {
			j -= 2
			b = toInt(input[j])
		}
	}
	fmt.Println("Part 1:", total)
}

type Memory struct {
	pos, size int
}

func part2() {
	input := utils.ReadFile("input.txt")
	files, free, pos := []Memory{}, []Memory{}, 0
	for i := range input { // Parse input
		size := toInt(input[i])
		if i%2 == 0 {
			files = append(files, Memory{pos, size})
		} else {
			free = append(free, Memory{pos, size})
		}
		pos += size
	}
	for i := len(files) - 1; i >= 0; i-- { // Move files
		file := files[i]
		for j, space := range free {
			if space.pos > file.pos {
				break
			}
			if space.size >= file.size {
				files[i] = Memory{space.pos, file.size}
				free[j] = Memory{space.pos + file.size, space.size - file.size}
				break
			}
		}
	}
	total := int64(0)
	for i, file := range files { // Calculate checksum
		for j := 0; j < file.size; j++ {
			total += int64((file.pos + j) * i)
		}
	}
	fmt.Println("Part 2:", total)
}

func main() {
	utils.TimeFunction(part1)
	utils.TimeFunction(part2)
}
