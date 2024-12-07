package main

import (
	"aoc2024/lib"
	"fmt"
	"math"
	"strings"
)

type operation func(int, int) int

func main() {
	equations := lib.ReadInputAsLines()

	calibrationResultPart1 := 0
	calibrationResultPart2 := 0
	for _, equation := range equations {
		parts := strings.Split(equation, ": ")
		result := lib.MustParseToInt(parts[0])
		var inputs []int
		for _, input := range strings.Split(parts[1], " ") {
			inputs = append(inputs, lib.MustParseToInt(input))
		}

		if canBeSolved(result, inputs, []operation{add, multiply}) {
			calibrationResultPart1 += result
		}

		if canBeSolved(result, inputs, []operation{add, multiply, concat}) {
			calibrationResultPart2 += result
		}
	}

	fmt.Printf("Part 1: %d\n", calibrationResultPart1)
	fmt.Printf("Part 2: %d\n", calibrationResultPart2)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	return a*int(math.Pow(10, digitsInNum(b))) + b
}

func digitsInNum(num int) float64 {
	// log10 only works from 2 onwards, assume we need to move at least one position
	if num < 2 {
		return 1
	}
	// add one to push 10 and 100 over to the next level
	return math.Ceil(math.Log10(float64(num + 1)))
}

func canBeSolved(result int, inputs []int, operations []operation) bool {
	return attemptToSolveRecurse(result, 0, inputs, operations)
}

func attemptToSolveRecurse(expectedResult, currentResult int, remainingInputs []int, operations []operation) bool {
	if remainingInputs == nil || len(remainingInputs) == 0 {
		return expectedResult == currentResult
	}

	for _, op := range operations {
		if attemptToSolveRecurse(expectedResult, op(currentResult, remainingInputs[0]), remainingInputs[1:], operations) {
			return true
		}
	}

	return false
}
