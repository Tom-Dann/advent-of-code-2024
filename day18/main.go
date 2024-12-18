package main

import (
	"container/heap"
	"fmt"
	"utils"
)

type Point struct{ x, y int }
type Item struct {
	pos         Point
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

func getNeighbours(p Point) []Point {
	return []Point{{p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}, {p.x, p.y - 1}}
}

func dijkstra(lines []string, maxTime int) int {
	size := 70
	var pq PriorityQueue
	points, times := map[Point]*Item{}, map[Point]int{}
	for i := 0; i <= size; i++ {
		for j := 0; j <= size; j++ {
			pos := Point{i, j}
			item := &Item{pos, 1e9, -1}
			points[pos] = item
			pq.Push(item)
		}
	}
	for t, line := range lines {
		var i, j int
		fmt.Sscanf(line, "%d,%d", &i, &j)
		times[Point{i, j}] = t + 1
	}
	heap.Init(&pq)
	pq.update(points[Point{0, 0}], 0)
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item)
		if u.pos.x == size && u.pos.y == size {
			return u.dist
		}
		for _, next := range getNeighbours(u.pos) {
			v := points[next]
			if v != nil && (times[v.pos] == 0 || times[v.pos] > maxTime) {
				if alt := u.dist + 1; alt < v.dist {
					pq.update(v, alt)
				}
			}
		}
	}
	return 1e9
}

func part1(lines []string) {
	fmt.Println("Part 1:", dijkstra(lines, 1024))
}

func part2(lines []string) {
	min, max := 1024, len(lines)
	for min < max-1 { // Binary search
		mid := (max + min) >> 1
		if dijkstra(lines, mid) == 1e9 {
			max = mid
		} else {
			min = mid
		}
	}
	fmt.Println("Part 2:", lines[min])
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, input)
	utils.TimeFunctionInput(part2, input)
}
