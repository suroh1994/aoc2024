package main

import (
	"aoc2024/lib"
	"fmt"
	"regexp"
)

var (
	regex         = regexp.MustCompile(`(\d+), Y.(\d+)$`)
	aCostPerPress = 3
	bCostPerPress = 1
	maxPresses    = 100
)

func main() {
	input := lib.ReadInputAsLines()
	machineCount := (len(input) + 1) / 4
	totalTokensUsed := 0
	for i := 0; i < machineCount; i++ {
		machineIdx := i * 4
		aMatches := regex.FindAllStringSubmatch(input[machineIdx], -1)[0][1:]
		bMatches := regex.FindAllStringSubmatch(input[machineIdx+1], -1)[0][1:]
		prizeMatches := regex.FindAllStringSubmatch(input[machineIdx+2], -1)[0][1:]

		targetX := lib.MustParseToInt(prizeMatches[0])
		targetY := lib.MustParseToInt(prizeMatches[1])
		aXPerPress := lib.MustParseToInt(aMatches[0])
		aYPerPress := lib.MustParseToInt(aMatches[1])
		bXPerPress := lib.MustParseToInt(bMatches[0])
		bYPerPress := lib.MustParseToInt(bMatches[1])

		bPressesToGetCloseToX := targetX / bXPerPress
		bPressesToGetCloseToY := targetY / bYPerPress

		maxBPresses := min(maxPresses, min(bPressesToGetCloseToX, bPressesToGetCloseToY))

		aPresses, bPresses := minButtonPresses(targetX, targetY, maxBPresses, aXPerPress, aYPerPress, bXPerPress, bYPerPress)
		if aPresses == -1 {
			continue
		}
		totalTokensUsed += aPresses*aCostPerPress + bPresses*bCostPerPress
	}

	fmt.Println("Part 1:", totalTokensUsed)
}

func minButtonPresses(targetX, targetY, maxBPresses, aXPerPress, aYPerPress, bXPerPress, bYPerPress int) (int, int) {
	for j := maxBPresses; j > 0; j-- {
		xDistanceLeft := targetX - j*bXPerPress
		yDistanceLeft := targetY - j*bYPerPress
		aPressesToGetCloseToX := xDistanceLeft / aXPerPress
		aPressesToGetCloseToY := yDistanceLeft / aYPerPress

		if aPressesToGetCloseToX == aPressesToGetCloseToY &&
			xDistanceLeft%aXPerPress == 0 && yDistanceLeft%aYPerPress == 0 &&
			aPressesToGetCloseToX <= 100 {
			return aPressesToGetCloseToX, j
		}
	}
	return -1, -1
}
