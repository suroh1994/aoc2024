package main

import (
	"aoc2024/lib"
	"fmt"
)

func main() {
	runeMap := lib.ReadInputAsRuneMap()

	wordCount := 0
	for x := range runeMap {
		for y := range runeMap[x] {
			if runeMap[x][y] == 'X' {
				wordCount += searchNeigbours(runeMap, x, y)
			}
		}
	}
	fmt.Printf("Part 1: %d\n", wordCount)
}

func searchNeigbours(runeMap [][]rune, x, y int) int {
	wordsFound := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !lib.IsInBounds(runeMap, x+i, y+j) {
				continue
			}

			if runeMap[x+i][y+j] == 'M' && lib.IsInBounds(runeMap, x+3*i, y+3*j) {
				if followDirection(runeMap, x, y, i, j) {
					wordsFound++
				}
			}
		}
	}
	return wordsFound
}

func followDirection(runeMap [][]rune, x, y, i, j int) bool {
	return runeMap[x+2*i][y+2*j] == 'A' && runeMap[x+3*i][y+3*j] == 'S'
}
