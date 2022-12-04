package main

import (
	"aoc/futils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := futils.Lines("../december-4-2022/input")
	fmt.Printf("%d fully contained\n", Overlapped(lines, true))
	fmt.Printf("%d overlapping\n", Overlapped(lines, false))
}

func Overlapped(lines []string, completely bool) int {
	ct := 0
	for _, l := range lines {
		a, b, x, y := Parse(l)
		if Contained(completely, a, b, x, y) {
			ct = ct + 1
		}
	}
	return ct
}

func Parse(in string) (int, int, int, int) {
	provided := []string{}
	ranges := []int{}
	for _, p := range strings.Split(in, ",") {
		provided = append(provided, strings.Split(p, "-")...)
	}
	for _, r := range provided {
		item, _ := strconv.Atoi(r)
		ranges = append(ranges, item)
	}
	return ranges[0], ranges[1], ranges[2], ranges[3]
}

func Contained(completely bool, a, b, x, y int) bool {
	switch {
	case completely && fullyContained(a, b, x, y):
		return true
	case !completely && partiallyContained(a, b, x, y):
		return true
	}
	return false
}

func fullyContained(a, b, x, y int) bool {
	return inside(a, b, x, y) || inside(x, y, a, b)
}

func inside(a, b, x, y int) bool {
	return (a >= x && b <= y)
}

func partiallyContained(a, b, x, y int) bool {
	return between(a, b, x, y) || between(x, y, a, b)
}

func between(a, b, x, y int) bool {
	return (a >= x && a <= y) || (b >= x && b <= y)
}
