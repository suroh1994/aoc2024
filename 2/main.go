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
		if areLevelsSafe(levels, 0) {
			safeLevels++
		}
	}

	fmt.Printf("Part 1: %d\n", safeLevels)
	/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
	safeLevels = 0
	for _, levels := range data {
		if areLevelsSafe(levels, 1) {
			safeLevels++
		}
	}

	fmt.Printf("Part 2: %d\n", safeLevels)
}

func areLevelsSafe(levels []int, numTolerableErrors int) bool {
	expectedSign := math.Signbit(float64(levels[1] - levels[0]))
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		absDiff := math.Abs(float64(diff))
		if absDiff > 3 || absDiff < 1 || math.Signbit(float64(diff)) != expectedSign {
			numTolerableErrors--
			if numTolerableErrors < 0 {
				return false
			} else {
				// Hack because if the first digit is the problem, we will only notice on the second digit
				levelsWithoutFirst := make([]int, len(levels)-1)
				copy(levelsWithoutFirst, levels[1:])

				levelsWithoutI := make([]int, len(levels)-1)
				copy(levelsWithoutI, levels[:i])
				copy(levelsWithoutI[i:], levels[i+1:])

				levelsWithoutIPlus1 := make([]int, len(levels)-1)
				copy(levelsWithoutIPlus1, levels[:i+1])
				copy(levelsWithoutIPlus1[i+1:], levels[i+2:])

				return areLevelsSafe(levelsWithoutFirst, numTolerableErrors-1) ||
					areLevelsSafe(levelsWithoutI, numTolerableErrors-1) ||
					areLevelsSafe(levelsWithoutIPlus1, numTolerableErrors-1)
			}
		}
	}

	return true
}
