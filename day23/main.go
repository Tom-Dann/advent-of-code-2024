package main

import (
	"fmt"
	"sort"
	"strings"
	"utils"
)

type StringSet map[string]bool

func lanKey(arr []string) string {
	sort.Strings(arr)
	return strings.Join(arr, ",")
}

func solve(lines []string) {
	graph := map[string]StringSet{}
	seen := map[string]bool{}
	for _, line := range lines {
		s := strings.Split(line, "-")
		for _, node := range []string{s[0], s[1]} {
			if !seen[node] {
				graph[node] = StringSet{}
				seen[node] = true
			}
		}
		graph[s[0]][s[1]] = true
		graph[s[1]][s[0]] = true
	}
	groups := map[string]struct{}{}
	triplets := map[string]struct{}{}
	for node, connections := range graph {
		for a := range connections {
			for b := range graph[a] {
				if connections[b] {
					key := lanKey([]string{node, a, b})
					groups[key] = struct{}{}
					if node[0] == 't' {
						triplets[key] = struct{}{}
					}
				}
			}
		}
	}
	for len(groups) > 2 {
		newGroups := map[string]struct{}{}
		for key := range groups {
			nodes := strings.Split(key, ",")
			for new := range seen {
				connected := true
				for _, node := range nodes {
					if node == new || !graph[node][new] {
						connected = false
						break
					}
				}
				if connected {
					newGroups[lanKey(append(nodes, new))] = struct{}{}
				}
			}
		}
		groups = newGroups
	}
	password := ""
	for k := range groups {
		password = k
	}
	fmt.Println("Part 1:", len(triplets))
	fmt.Println("Part 2:", password)
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n"))
}
