package main

import (
	"testing"
)

func testData() []string {
	return []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
}

func TestVisibleTrees(t *testing.T) {
	lines := testData()
	expect := 21
	got := VisibleTrees(lines)
	if got != expect {
		t.Errorf("Expected %d, go %d", expect, got)
	}
}
