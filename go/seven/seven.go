package main

import (
	"aoc/futils"
	"aoc/seven/node"
	"fmt"
	"strconv"
	"strings"
)

const (
	totalDiskSpace    int = 70000000
	requiredDiskSpace int = 30000000
)

func main() {
	lines := futils.Lines("../december-7-2022/input")
	fmt.Printf("Part One: %d\n", ParseShellHistory(lines).SizeFilter(100000, 0))
	fmt.Printf("Part Two: %s\n", ParseShellHistory(lines).FindSpace(totalDiskSpace, requiredDiskSpace).Path())
}

func ParseShellHistory(lines []string) *node.Node {
	root := &node.Node{Type: node.Directory}
	n := root
	for _, line := range lines {
		switch line[:4] {
		case "$ cd":
			n = cd(n, line)
		case "$ ls":
			n = ls(n, line)
		case "dir ":
			n = dir(n, line)
		default:
			n = file(n, line)
		}
	}

	root.DirSizes()
	return root
}

func cd(n *node.Node, line string) *node.Node {
	l := strings.Split(line, " ")
	switch l[2] {
	case "/":
		n = n.Root()
	case "..":
		n = n.Parent()
	default:
		n = n.FindOrCreateChild(node.Directory, l[2], 0)
	}
	return n
}

func ls(n *node.Node, line string) *node.Node {
	return n
}

func dir(n *node.Node, line string) *node.Node {
	l := strings.Split(line, " ")
	n.FindOrCreateChild(node.Directory, l[1], 0)
	return n
}

func file(n *node.Node, line string) *node.Node {
	l := strings.Split(line, " ")
	size, _ := strconv.Atoi(l[0])
	n.FindOrCreateChild(node.File, l[1], size)
	return n
}
