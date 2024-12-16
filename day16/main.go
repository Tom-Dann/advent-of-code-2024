package main

import (
	"container/heap"
	"fmt"
	"utils"
)

type Point struct{ x, y int }
type State struct{ pos, dir Point }
type Item struct {
	value       State
	dist, index int
}
type PriorityQueue []*Item

// Priority Queue implementation https://pkg.go.dev/container/heap
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].index, pq[j].index = i, j }
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) update(item *Item, dist int) { item.dist = dist; heap.Fix(pq, item.index) }

func dijkstra(states map[State]*Item, pq PriorityQueue, start, target Point) (int, int) {
	heap.Init(&pq)
	pq.update(states[State{start, Point{1, 0}}], 0)
	prev := map[*Item][]*Item{}
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item)
		if u.value.pos == target {
			visited := map[Point]struct{}{}
			queue := []*Item{u}
			for len(queue) > 0 {
				next := queue[0]
				queue = queue[1:]
				visited[next.value.pos] = struct{}{}
				queue = append(queue, prev[next]...)
			}
			return u.dist, len(visited)
		}
		for _, dir := range []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			pos := u.value.pos
			dist := 1000
			if u.value.dir == dir {
				pos.x += dir.x
				pos.y += dir.y
				dist = 1
			}
			v := states[State{pos, dir}]
			if v != nil {
				if alt := u.dist + dist; alt <= v.dist {
					if alt < v.dist {
						pq.update(v, alt)
						prev[v] = []*Item{u}
					} else {
						prev[v] = append(prev[v], u)
					}
				}
			}
		}
	}
	return -1, -1
}

func solve(lines []string) {
	grid := map[Point]rune{}
	var target, start Point
	var pq PriorityQueue
	states := map[State]*Item{}
	for j, line := range lines {
		for i, char := range line {
			grid[Point{i, j}] = char
			if char != '#' {
				for _, d := range []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					state := State{Point{i, j}, d}
					item := &Item{state, 1e9, -1}
					states[state] = item
					pq.Push(item)
				}
				if char == 'E' {
					target = Point{i, j}
				} else if char == 'S' {
					start = Point{i, j}
				}
			}
		}
	}
	part1, part2 := dijkstra(states, pq, start, target)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n"))
}
