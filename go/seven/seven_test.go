package main

import (
	"testing"
)

func testData() []string {
	return []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}
}

func TestParseShellHistory(t *testing.T) {
	lines := testData()
	expect := "/d/k (7214296 bytes)"
	root := ParseShellHistory(lines)
	got := root.Children()[3].Children()[3].Path()
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}

func TestNodeSizes(t *testing.T) {
	lines := testData()
	root := ParseShellHistory(lines)
	if root.Size != 48381165 {
		t.Errorf("Expected root size of 48381165, got %d", root.Size)
	}
}

func TestSizeFilter(t *testing.T) {
	lines := testData()
	root := ParseShellHistory(lines)
	expect := 95437
	got := root.SizeFilter(100000, 0)
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestFindSpace(t *testing.T) {
	lines := testData()
	root := ParseShellHistory(lines)
	// Given total space of 60,000,000 and required space of 25,000,000
	// 11,619,415 - (60,000,000 - 48,381,165) = 580 free space needed
	expect := 584 // directory node larger but closest in size
	got := root.FindSpace(60000000, 11619415)
	if expect != got.Size {
		t.Errorf("Expected %d, got %d", expect, got.Size)
	}
}
