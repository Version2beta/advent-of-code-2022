package main

import (
	"testing"
)

func testData() ([]string, []int) {
	signals := []string{
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
	startPos := []int{5, 6, 10, 11}
	return signals, startPos
}

func TestFrame(t *testing.T) {
	signals, startPos := testData()
	for i, s := range signals {
		expect := startPos[i]
		got := StartOfPacket(s)
		if got != expect {
			t.Errorf("Expected %d, got %d", expect, got)
		}
	}
}
