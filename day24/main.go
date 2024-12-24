package main

import (
	"fmt"
	"sort"
	"strings"
	"utils"
)

const zSize = 45

type Gate struct{ a, op, b, out string }

func logic(op string, a, b bool) bool {
	switch op {
	case "AND":
		return a && b
	case "OR":
		return a || b
	case "XOR":
		return a != b
	}
	panic(fmt.Sprint("Unknown operation", op))
}

func output(wires map[string]bool, queue []Gate) int {
	var gate Gate
	for len(queue) > 0 {
		gate, queue = queue[0], queue[1:]
		a, aFound := wires[gate.a]
		b, bFound := wires[gate.b]
		if aFound && bFound {
			wires[gate.out] = logic(gate.op, a, b)
		} else {
			queue = append(queue, gate)
		}
	}
	z, mult := 0, 1
	for i := 0; i < zSize; i++ {
		wire := fmt.Sprintf("z%02d", i)
		if wires[wire] {
			z += mult
		}
		mult <<= 1
	}
	return z
}

func getInputs(x, y int) map[string]bool {
	wires := map[string]bool{}
	for i := 0; i < zSize; i++ {
		wires[fmt.Sprintf("x%02d", i)] = x>>i%2 == 1
		wires[fmt.Sprintf("y%02d", i)] = y>>i%2 == 1
	}
	return wires
}

func printGraphviz(lines []string) {
	fmt.Println("graph network {")
	for i, line := range lines {
		var a, op, b, out string
		fmt.Sscanf(line, "%s %s %s -> %s", &a, &op, &b, &out)
		fmt.Printf("	%00d[label=%s];\n", i, op)
		fmt.Printf("	%00d -- { %s %s };\n", i, a, b)
		fmt.Printf("	%s -- { %00d };\n", out, i)
	}
	fmt.Println("}")
}

func countErrors(gates []Gate) int {
	errors := 0
	for i := 0; i < zSize-1; i++ {
		test := 1 << i
		z := output(getInputs(test, 0), gates)
		if z != test {
			errors++
			fmt.Println("Error found at index", i)
		}
	}
	return errors
}

func parseWires(lines []string) map[string]bool {
	wires := map[string]bool{}
	for _, line := range lines {
		s := strings.Split(line, ": ")
		wires[s[0]] = s[1] == "1"
	}
	return wires
}

func parseGates(lines []string) []Gate {
	gates := []Gate{}
	for _, line := range lines {
		var a, op, b, out string
		fmt.Sscanf(line, "%s %s %s -> %s", &a, &op, &b, &out)
		gates = append(gates, Gate{a, op, b, out})
	}
	return gates
}

func solve(input []string) {
	wires := parseWires(strings.Split(input[0], "\n"))
	gates := parseGates(strings.Split(input[1], "\n"))
	fmt.Println("Part 1:", output(wires, gates))

	fmt.Println("Pre swap error count:", countErrors(gates))
	type Swap struct{ op, old, new string }
	swaps := []Swap{
		{"XOR", "rts", "z07"}, {"OR", "z07", "rts"},
		{"XOR", "jpj", "z12"}, {"AND", "z12", "jpj"},
		{"XOR", "kgj", "z26"}, {"AND", "z26", "kgj"},
		{"AND", "chv", "vvw"}, {"XOR", "vvw", "chv"},
	}
	for i, gate := range gates {
		for _, swap := range swaps {
			if gate.op == swap.op && gate.out == swap.old {
				gates[i].out = swap.new
			}
		}
	}
	fmt.Println("Post swap error count:", countErrors(gates))

	swappedWires := []string{}
	for _, swap := range swaps {
		swappedWires = append(swappedWires, swap.old)
	}
	sort.Strings(swappedWires)
	fmt.Println("Part 2:", strings.Join(swappedWires, ","))
}

func main() {
	utils.TimeFunctionInput(solve, utils.ReadInput("input.txt", "\n\n"))
}
