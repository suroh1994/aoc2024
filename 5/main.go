package main

import (
	"aoc2024/lib"
	"fmt"
	"slices"
	"strings"
)

func main() {
	input := lib.ReadInputAsLines()
	rules, manuals := parseInput(input)

	rulesMap := rules2Map(rules)

	midPageSum := 0
	for _, manual := range manuals {
		if isPrintedAccordingToRules(manual, rulesMap) {
			midPageSum += manual[len(manual)/2]
		}
	}

	fmt.Printf("Part 1: %d\n", midPageSum)
	/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
	midPageSum = 0
	rulesBasedSort := generateSortFunction(rulesMap)
	for _, manual := range manuals {
		if !isPrintedAccordingToRules(manual, rulesMap) {
			slices.SortFunc(manual, rulesBasedSort)
			midPageSum += manual[len(manual)/2]
		}
	}
	fmt.Printf("Part 2: %d\n", midPageSum)
}
func parseInput(input []string) ([]string, [][]int) {
	rules := []string{}
	manuals := [][]int{}

	parsingRules := true
	for _, line := range input {
		if line == "" {
			parsingRules = false
			continue
		}

		if parsingRules {
			rules = append(rules, line)
		} else {
			manual := []int{}
			for _, numberStr := range strings.Split(line, ",") {
				manual = append(manual, lib.MustParseToInt(numberStr))
			}
			manuals = append(manuals, manual)
		}
	}

	return rules, manuals
}

func rules2Map(rules []string) map[int][]int {
	rulesMap := make(map[int][]int)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		firstPage := lib.MustParseToInt(parts[0])
		secondPage := lib.MustParseToInt(parts[1])
		rulesMap[firstPage] = append(rulesMap[firstPage], secondPage)
	}
	return rulesMap
}

func isPrintedAccordingToRules(manual []int, rulesMap map[int][]int) bool {
	for idx := len(manual) - 1; idx > 0; idx-- {
		pagesToBePrintedAfterCurrentPage := rulesMap[manual[idx]]
		for _, earlierPage := range manual[:idx] {
			if slices.Contains(pagesToBePrintedAfterCurrentPage, earlierPage) {
				return false
			}
		}
	}
	return true
}

func generateSortFunction(rulesMap map[int][]int) func(int, int) int {
	return func(pageA int, pageB int) int {
		rulesPageA := rulesMap[pageA]
		rulesPageB := rulesMap[pageB]
		if slices.Contains(rulesPageA, pageB) {
			return -1
		}
		if slices.Contains(rulesPageB, pageA) {
			return 1
		}
		return 0
	}
}
