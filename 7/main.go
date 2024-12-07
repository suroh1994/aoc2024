package main

import (
	"aoc2024/lib"
	"fmt"
	"strings"
)

var (
	operations = []func(int, int) int{add, multiply}
)

func main() {
	equations := lib.ReadInputAsLines()

	calibrationResult := 0
	for _, equation := range equations {
		parts := strings.Split(equation, ": ")
		result := lib.MustParseToInt(parts[0])
		var inputs []int
		for _, input := range strings.Split(parts[1], " ") {
			inputs = append(inputs, lib.MustParseToInt(input))
		}

		if canBeSolved(result, inputs) {
			calibrationResult += result
		}
	}

	fmt.Printf("Part 1: %d\n", calibrationResult)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func canBeSolved(result int, inputs []int) bool {
	return attemptToSolveRecurse(result, 0, inputs)
}

func attemptToSolveRecurse(expectedResult, currentResult int, remainingInputs []int) bool {
	if remainingInputs == nil || len(remainingInputs) == 0 {
		return expectedResult == currentResult
	}

	for _, operation := range operations {
		if attemptToSolveRecurse(expectedResult, operation(currentResult, remainingInputs[0]), remainingInputs[1:]) {
			return true
		}
	}

	return false
}
