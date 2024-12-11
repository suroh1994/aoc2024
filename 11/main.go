package main

import (
	"aoc2024/lib"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := lib.ReadInput()
	numberStrings := strings.Split(input, " ")

	numbers := make([]int, len(numberStrings), 10000000) // let's start higher...
	for idx := range numberStrings {
		numbers[idx] = lib.MustParseToInt(numberStrings[idx])
	}

	stoneCount := 0
	for _, num := range numbers {
		stoneCount += blinkAtStone(num, 25)
	}

	fmt.Printf("Part 1: %d\n", stoneCount)
	/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
	stoneCount = 0
	for _, num := range numbers {
		stoneCount += blinkAtStone(num, 75)
	}
	fmt.Printf("Part 2: %d\n", stoneCount)
}

var (
	cache = map[lib.Point2D]int{}
)

func blinkAtStone(stoneValue, blinksLeft int) int {
	if stoneCount, exists := cache[lib.NewPoint2D(stoneValue, blinksLeft)]; exists {
		return stoneCount
	}

	if blinksLeft == 0 {
		cache[lib.NewPoint2D(stoneValue, blinksLeft)] = 1
		return 1
	}

	if stoneValue == 0 {
		stoneCount := blinkAtStone(1, blinksLeft-1)
		cache[lib.NewPoint2D(stoneValue, blinksLeft)] = stoneCount
		return stoneCount
	}

	digitCount := lib.DigitsInNum(stoneValue)
	if digitCount%2 == 0 {
		left := stoneValue / int(math.Pow(10, float64(lib.DigitsInNum(stoneValue)/2)))
		right := stoneValue % int(math.Pow(10, float64(lib.DigitsInNum(stoneValue)/2)))
		stoneCount := blinkAtStone(left, blinksLeft-1) + blinkAtStone(right, blinksLeft-1)
		cache[lib.NewPoint2D(stoneValue, blinksLeft)] = stoneCount
		return stoneCount
	}

	stoneCount := blinkAtStone(stoneValue*2024, blinksLeft-1)
	cache[lib.NewPoint2D(stoneValue, blinksLeft)] = stoneCount
	return stoneCount
}
