package main

import (
	"aoc/futils"
	"fmt"
)

func main() {
	lines := futils.Lines("../december-3-2022/input")
	fmt.Printf("Priorities part one: %d\n", Priorities(lines))
	fmt.Printf("Priorities part two: %d\n", GroupPriorities(Group(lines)))
}

func Priorities(lines []string) int {
	sum := 0
	for _, l := range lines {
		sum = sum + Score(Intersection(split(l)...))
	}
	return sum
}

func Group(group []string) [][]string {
	groups := [][]string{}
	for { // why yes, this _is_ a reducer without recursion
		if len(group) == 0 {
			break
		}

		size := min(3, len(group))
		groups = append(groups, group[:size])
		group = group[size:]
	}
	return groups
}

func GroupPriorities(groups [][]string) int {
	sum := 0
	for _, g := range groups {
		sum = sum + Score(Intersection(g...))
	}
	return sum
}

func Intersection(inputs ...string) string {
	if len(inputs) < 2 {
		return ""
	}

	for { // why yes, this _is_ also a reducer without recursion
		if len(inputs) == 1 {
			break
		}

		hash := map[rune]struct{}{}
		intersections := map[rune]struct{}{}
		l, r := inputs[0], inputs[1]

		for _, v := range l {
			hash[v] = struct{}{}
		}
		for _, v := range r {
			if _, ok := hash[v]; ok {
				intersections[v] = struct{}{}
			}
		}

		intersection := ""
		for k := range intersections {
			intersection = intersection + string(k)
		}
		inputs = append([]string{intersection}, inputs[2:]...)
	}

	return inputs[0]
}

func Score(input string) int {
	sum := 0
	for _, c := range input {
		switch {
		case c >= 'A' && c <= 'Z':
			sum = sum + int(c) - 'A' + 26 + 1
		case c >= 'a' && c <= 'z':
			sum = sum + int(c) - 'a' + 1
		}
	}

	return sum
}

func split(line string) []string {
	split := len(line) / 2
	return []string{line[:split], line[split:]}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
