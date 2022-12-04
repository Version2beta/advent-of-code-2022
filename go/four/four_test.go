package main

import (
	"testing"
)

func testData() []string {
	return []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9", // overlapped
		"2-8,3-7", // contained, overlapped
		"6-6,4-6", // contained, overlapped
		"2-6,4-8", // overlapped
	}
}

func TestFullyOverlapped(t *testing.T) {
	expect := 2
	got := Overlapped(testData(), true)
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestPartiallyOverlapped(t *testing.T) {
	expect := 4
	got := Overlapped(testData(), false)
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestParse(t *testing.T) {
	expect := [][]int{
		{2, 4, 6, 8},
		{2, 3, 4, 5},
		{5, 7, 7, 9},
		{2, 8, 3, 7},
		{6, 6, 4, 6},
		{2, 6, 4, 8},
	}
	for i, in := range testData() {
		a, b, x, y := Parse(in)
		if a != expect[i][0] || b != expect[i][1] || x != expect[i][2] || y != expect[i][3] {
			t.Errorf("Expected %#v, got %#v", expect[i], []int{a, b, x, y})
		}
	}
}

func TestFullyContained(t *testing.T) {
	expect := []bool{false, false, false, true, true, false}
	for i, l := range testData() {
		a, b, x, y := Parse(l)
		got := Contained(true, a, b, x, y)
		if got != expect[i] {
			t.Errorf("Expected %t, got %t in loop %d", expect[i], got, i)
		}
	}
}

func TestPartiallyContained(t *testing.T) {
	expect := []bool{false, false, true, true, true, true}
	for i, l := range testData() {
		a, b, x, y := Parse(l)
		got := Contained(false, a, b, x, y)
		if got != expect[i] {
			t.Errorf("Expected %t, got %t in loop %d", expect[i], got, i)
		}
	}
}
