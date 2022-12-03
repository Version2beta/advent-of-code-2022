package main

import (
	"testing"
)

func testData() []string {
	return []string{
		"A Y",
		"B X",
		"C Z",
	}
}

func TestRockPaperScissors(t *testing.T) {
	expect := 15
	got := RockPaperScissors(testData())
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestStrategicRockPaperScissors(t *testing.T) {
	expect := 12
	got := StrategicRockPaperScissors(testData())
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}
