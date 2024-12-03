package main

import (
	"aoc2024/lib"
	"fmt"
	"regexp"
)

func main() {
	content := lib.ReadInput()
	totalPart1, totalPart2 := calculateInstructions(content)

	fmt.Printf("Part 1: %d\n", totalPart1)
	fmt.Printf("Part 2: %d\n", totalPart2)
}

var regexWithToggles = regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do(n't)*\\(\\)")

func calculateInstructions(line string) (int, int) {
	matches := regexWithToggles.FindAllStringSubmatch(line, -1)
	total := 0
	toggle := true
	totalWithToggle := 0
	for _, match := range matches {
		switch match[0] {
		case "do()":
			toggle = true
		case "don't()":
			toggle = false
		default:
			total += lib.MustParseToInt(match[1]) * lib.MustParseToInt(match[2])
			if toggle {
				totalWithToggle += lib.MustParseToInt(match[1]) * lib.MustParseToInt(match[2])
			}
		}
	}
	return total, totalWithToggle
}
