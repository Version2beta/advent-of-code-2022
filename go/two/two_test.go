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

func TestRPS(t *testing.T) {
	expect := 15
	got := RPS(testData())
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestStrategicRPS(t *testing.T) {
	expect := 12
	got := StrategicRPS(testData())
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}
