package main

import (
	"aoc/five/stack"
	"aoc/futils"
	"fmt"
	"strconv"
	"strings"
)

// [N]             [R]             [C]
// [T] [J]         [S] [J]         [N]
// [B] [Z]     [H] [M] [Z]         [D]
// [S] [P]     [G] [L] [H] [Z]     [T]
// [Q] [D]     [F] [D] [V] [L] [S] [M]
// [H] [F] [V] [J] [C] [W] [P] [W] [L]
// [G] [S] [H] [Z] [Z] [T] [F] [V] [H]
// [R] [H] [Z] [M] [T] [M] [T] [Q] [W]
//  1   2   3   4   5   6   7   8   9

func startData() []*stack.Stack {
	return []*stack.Stack{
		{Items: []rune("NTBSQHGR")},
		{Items: []rune("JZPDFSH")},
		{Items: []rune("VHZ")},
		{Items: []rune("HGFJZM")},
		{Items: []rune("RSMLDCZT")},
		{Items: []rune("JZHVWTM")},
		{Items: []rune("ZLPFT")},
		{Items: []rune("SWVQ")},
		{Items: []rune("CNDTMLHW")},
	}
}

func main() {
	stacks := startData()
	lines := futils.Lines("../december-5-2022/input")
	partOne := TopCrates(PartOne(lines, stacks))

	stacks = startData()
	partTwo := TopCrates(PartTwo(lines, stacks))

	fmt.Printf("Part one: %s\n", partOne)
	fmt.Printf("Part two: %s\n", partTwo)
}

func TopCrates(stacks []*stack.Stack) string {
	tops := &stack.Stack{}
	ct := len(stacks) - 1
	for i := range stacks {
		tops.Push(stacks[ct-i].Items[:1])
	}
	return tops.String()
}

func PartOne(lines []string, stacks []*stack.Stack) []*stack.Stack {
	for i, line := range lines {
		fmt.Printf("Row %d: ", i)
		n, from, to := stacksFromLine(stacks, line)
		stack.PopPush(n, from, to)
	}

	return stacks
}

func PartTwo(lines []string, stacks []*stack.Stack) []*stack.Stack {
	for i, line := range lines {
		fmt.Printf("Row %d: ", i)
		n, from, to := stacksFromLine(stacks, line)
		stack.Move(n, from, to)
	}

	return stacks
}

func stacksFromLine(stacks []*stack.Stack, line string) (int, *stack.Stack, *stack.Stack) {
	mv := strings.Split(line, ",")

	n, err := strconv.Atoi(mv[0])
	futils.CheckErr(err)
	fromNum, err := strconv.Atoi(mv[1])
	futils.CheckErr(err)
	toNum, err := strconv.Atoi(mv[2])
	futils.CheckErr(err)

	from, to := stacks[fromNum-1], stacks[toNum-1]
	return n, from, to
}
