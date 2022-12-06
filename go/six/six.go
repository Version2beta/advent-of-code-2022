package main

import (
	"aoc/futils"
	"fmt"
)

func main() {
	lines := futils.Lines("../december-6-2022/input")
	fmt.Printf("Part one: %d\n", StartPos(lines, 4)[0])
	fmt.Printf("Part two: %d\n", StartPos(lines, 14)[0])
}

func StartPos(lines []string, l int) []int {
	res := []int{}
	for _, line := range lines {
		res = append(res, StartOfPacket(line, l))
	}
	return res
}

func StartOfPacket(signal string, l int) int {
	for i := l; i < len(signal); i++ {
		if distinct(signal[i-l : i]) {
			return i
		}
	}
	return 0 // not found
}

func distinct(s string) bool {
	set := map[rune]struct{}{}
	for _, c := range s {
		set[c] = struct{}{}
	}
	return len(set) == len(s)
}
