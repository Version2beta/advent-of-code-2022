package main

import (
	"testing"
)

func testData() []string {
	return []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",         // p
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", // L
		"PmmdzqPrVvPwwTWBwg",               // P
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",   // v
		"ttgJtRGJQctTZtZT",                 // t
		"CrZsJsPPZsGzwwsLwLmpwMDw",         // s
	}
}

func testGroupData() []string {
	return []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg", // group one, 'r'
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw", // group two, 'Z'
	}
}

func TestPriorities(t *testing.T) {
	expect := 157
	got := Priorities(testData())
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestIntersection(t *testing.T) {
	expect := "pLPvts"
	got := ""
	for _, s := range testData() {
		got = got + Intersection(split(s)...)
	}
	if got != expect {
		t.Errorf("Expected %s, got %s", expect, got)
	}
}

func TestScorer(t *testing.T) {
	expect := 106
	got := Score("AZaz")
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
	expect = 157
	got = Score("pLPvts")
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestGroup(t *testing.T) {
	data := testGroupData()
	expectGroupsLen, expectMembersLen := 2, 3
	got := Group(data)
	if len(got) != expectGroupsLen {
		t.Errorf("Expected %d, got %d", expectGroupsLen, len(got))
	}
	for i, g := range got {
		if len(g) != expectMembersLen {
			t.Errorf("Expected %d, got %d", expectMembersLen, len(g))
		}
		for j, m := range g {
			if m != data[i*3+j] {
				t.Errorf("Expected %s, got %s", data[i*3+j], m)
			}
		}
	}
}

func TestGroupPriorities(t *testing.T) {
	expect := 70
	got := GroupPriorities(Group(testGroupData()))
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}
