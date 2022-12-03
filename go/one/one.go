package main

import (
	"aoc/futils"
	"fmt"
	"strconv"
)

func main() {
	lines := futils.Lines("../december-1-2022/input")
	maxC, top3 := MaxC(lines), Top3(lines)
	fmt.Printf("Max calories for one elf: %d\n", maxC)
	fmt.Printf("Total calories from top 3 elves: %d\n", top3)
}

func MaxC(lines []string) int {
	maxC, perC := 0, 0
	for _, l := range lines {
		if l == "" {
			perC = 0
			continue
		}
		c, _ := strconv.Atoi(l)
		perC = perC + c
		if perC > maxC {
			maxC = perC
		}
	}
	return maxC
}

func Top3(lines []string) int {
	perC, top3 := 0, []int{0, 0, 0}
	for o, l := range lines {
		c, _ := strconv.Atoi(l)
		perC = perC + c

		if l == "" || o == len(lines)-1 {
			for i := range top3 {
				if top3[i] < perC {
					top3[i], perC = perC, top3[i]
				}
			}
			perC = 0
			continue
		}
	}
	return sliceSum(top3)
}

func sliceSum(in []int) int {
	sum := 0
	for _, v := range in {
		sum = sum + v
	}
	return sum
}
