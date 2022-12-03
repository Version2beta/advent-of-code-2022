package main

import (
	"testing"
)

func testData() []string {
	return []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
}

func TestMaxC(t *testing.T) {
	expect := 24000
	got := MaxC(testData())
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestTop3(t *testing.T) {
	expect := 45000
	got := Top3(testData())
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}
