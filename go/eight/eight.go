package main

import (
	"aoc/futils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := futils.Lines("../december-8-2022/input")
	fmt.Printf("Part one: %d\n", VisibleTrees(lines))
}

type Tree struct {
	height  int
	visible bool
}

type Forest [][]*Tree

func VisibleTrees(lines []string) int {
	var forest Forest
	forest = forest.parseTrees(lines)
	fmt.Println(forest)
	return 0
}

func (f *Forest) parseTrees(lines []string) Forest {
	var forest Forest
	var trees []*Tree

	for _, line := range lines {
		for _, h := range strings.Split(line, "") {
			height, _ := strconv.Atoi(h)
			tree := &Tree{
				height:  height,
				visible: false,
			}
			trees = append(trees, tree)
		}
		forest = append(forest, trees)
	}
	return forest
}

func (f *Forest) rotateTrees() {
	var rotated forest
	var trees []*Tree

	for _, outer := range f {
		for _, inner := range outer {

		}
	}
}
