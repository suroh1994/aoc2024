package main

import (
	"aoc2024/lib"
	"fmt"
	"strconv"
)

func main() {
	input := lib.ReadInput()
	expanded := expandInput(input)
	compress(expanded)

	checksum := 0
	for i := range expanded {
		if expanded[i] == -1 {
			break
		}
		checksum += i * expanded[i]
	}

	fmt.Println(checksum)
}

func expandInput(input string) []int {
	expandedInput := make([]int, 0)
	isEmpty := false
	for idx := range input {
		value := -1
		if !isEmpty {
			value = idx / 2
		}
		for count := 0; count < int(input[idx]-'0'); count++ {
			expandedInput = append(expandedInput, value)
		}

		isEmpty = !isEmpty
	}
	return expandedInput
}

func findNextEmpty(values []int) int {
	for i := range values {
		if values[i] == -1 {
			return i
		}
	}
	panic("no value found")
}

func findNextNotEmpty(values []int) int {
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] != -1 {
			return i
		}
	}
	panic("no non-empty value found")
}

func compress(input []int) {
	nextEmpty := findNextEmpty(input)
	nextNotEmpty := findNextNotEmpty(input)
	//viz(input)

	for nextEmpty < nextNotEmpty {
		input[nextEmpty] = input[nextNotEmpty]
		input[nextNotEmpty] = -1
		//viz(input)

		nextEmpty = findNextEmpty(input)
		nextNotEmpty = findNextNotEmpty(input)
	}
}

func viz(input []int) {
	for i := range input {
		symbol := strconv.Itoa(input[i])
		if input[i] == -1 {
			symbol = "."
		}
		fmt.Print(symbol)
	}
	fmt.Println()
}
