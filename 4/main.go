package main

import (
	"aoc2024/lib"
	"fmt"
	"slices"
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
	/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
	xMasCount := 0
	for x := range runeMap {
		for y := range runeMap[x] {
			if runeMap[x][y] == 'A' {
				if isXMas(runeMap, x, y) {
					xMasCount++
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", xMasCount)
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

func isXMas(runeMap [][]rune, x, y int) bool {
	if !lib.IsInBounds(runeMap, x-1, y-1) || !lib.IsInBounds(runeMap, x+1, y+1) {
		return false
	}

	words := []string{
		string([]rune{runeMap[x-1][y-1], runeMap[x][y], runeMap[x+1][y+1]}),
		string([]rune{runeMap[x-1][y+1], runeMap[x][y], runeMap[x+1][y-1]}),
	}

	return (words[0] == "MAS" || reverseString(words[0]) == "MAS") &&
		(words[0] == words[1] || words[0] == reverseString(words[1]))
}

func reverseString(s string) string {
	asRunes := []rune(s)
	slices.Reverse(asRunes)
	return string(asRunes)
}
