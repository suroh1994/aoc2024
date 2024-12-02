package main

import (
	"aoc2024/lib"
	"fmt"
	"math"
)

func main() {
	data := lib.ReadMultipleIntValuesPerLine(" ")

	safeLevels := 0
	for _, levels := range data {
		if areLevelsSafe(levels) {
			safeLevels++
		}
	}

	fmt.Printf("Part 1: %d\n", safeLevels)
}

func areLevelsSafe(levels []int) bool {
	isDescending := math.Signbit(float64(levels[1] - levels[0]))
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		absDiff := math.Abs(float64(diff))
		if absDiff > 3 || absDiff < 1 || math.Signbit(float64(diff)) != isDescending {
			fmt.Printf("%v is not safe because of %d and %d\n", levels, levels[i], levels[i+1])
			return false
		}
	}

	fmt.Printf("%v is safe\n", levels)
	return true
}
