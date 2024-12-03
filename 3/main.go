package main

import (
	"aoc2024/lib"
	"fmt"
	"regexp"
)

func main() {
	content := lib.ReadInputAsLines()
	total := 0
	for _, line := range content {
		total += calculateInstructions(line)
	}

	fmt.Printf("Part 1: %d\n", total)
}

var regex = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

func calculateInstructions(line string) int {
	matches := regex.FindAllStringSubmatch(line, -1)
	lineTotal := 0
	for _, match := range matches {
		lineTotal += lib.MustParseToInt(match[1]) * lib.MustParseToInt(match[2])
	}
	return lineTotal
}
