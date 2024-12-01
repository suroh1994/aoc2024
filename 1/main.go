package main

import (
	"aoc2024/lib"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := lib.ReadInputAsLines()

	listA := make([]int, len(lines))
	listB := make([]int, len(lines))
	for idx, line := range lines {
		numbers := strings.Split(line, "   ")
		listA[idx], _ = strconv.Atoi(numbers[0])
		listB[idx], _ = strconv.Atoi(numbers[1])
	}

	slices.Sort(listA)
	slices.Sort(listB)

	totalDist := 0
	for idx := range listA {
		totalDist += int(math.Abs(float64(listA[idx] - listB[idx])))
	}

	fmt.Println(totalDist)
}
