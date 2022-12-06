package main

import (
	"aoc/five/stack"
	"testing"
)

// [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2

func testData() ([]*stack.Stack, []string) {
	stacks := []*stack.Stack{
		{Items: []rune("NZ")},
		{Items: []rune("DCM")},
		{Items: []rune("P")},
	}
	return stacks, []string{
		"1,2,1",
		"3,1,3",
		"2,2,1",
		"1,1,2",
	}
}

func TestPartOne(t *testing.T) {
	stacks, lines := testData()
	expect := "CMZ"
	got := TopCrates(PartOne(lines, stacks))
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}

func TestPartTwo(t *testing.T) {
	stacks, lines := testData()
	expect := "MCD"
	got := TopCrates(PartTwo(lines, stacks))
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}
