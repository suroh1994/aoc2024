package main

import (
	"aoc2024/lib"
	"fmt"
	"regexp"
)

func main() {
	content := lib.ReadInputAsLines()
	totalPart1 := 0
	totalPart2 := 0
	toggle := true
	for _, line := range content {
		totalPart1 += calculateInstructions(line)
		totalPart2 += calculateInstructionsWithToggles(line, &toggle)
	}

	fmt.Printf("Part 1: %d\n", totalPart1)
	fmt.Printf("Part 2: %d\n", totalPart2)
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

var regexWithToggles = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do(n't)*\\(\\)")

func calculateInstructionsWithToggles(line string, toggle *bool) int {
	matches := regexWithToggles.FindAllStringSubmatch(line, -1)
	lineTotal := 0
	for _, match := range matches {
		switch match[0] {
		case "do()":
			*toggle = true
		case "don't()":
			*toggle = false
		default:
			if *toggle {
				lineTotal += lib.MustParseToInt(match[1]) * lib.MustParseToInt(match[2])
			}
		}
	}
	return lineTotal
}
