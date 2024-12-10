package main

import (
	"aoc2024/lib"
	"fmt"
	"maps"
)

var (
	directions = []lib.Point2D{
		lib.UP,
		lib.DOWN,
		lib.LEFT,
		lib.RIGHT,
	}
)

func main() {
	inputMap := lib.ReadInputAsRuneMap()
	// find all zeros
	zeroes := findAllZeroes(inputMap)

	trails := 0
	// for each zero
	for _, zero := range zeroes {
		//  find all tracks to 9s recursively
		trailHeadsReached := findAllNines(inputMap, zero)
		// count all 9s reached
		for range maps.Keys(trailHeadsReached) {
			// add to total
			trails++
		}
	}

	fmt.Printf("Part 1: %d\n", trails)
}

func findAllZeroes(inputMap [][]rune) []lib.Point2D {
	zeroes := make([]lib.Point2D, 0)
	for x := 0; x < len(inputMap); x++ {
		for y := 0; y < len(inputMap[x]); y++ {
			if inputMap[x][y] == '0' {
				zeroes = append(zeroes, lib.NewPoint2D(x, y))
			}
		}
	}
	return zeroes
}

func findAllNines(inputMap [][]rune, start lib.Point2D) map[lib.Point2D]bool {
	nines := make(map[lib.Point2D]bool)
	for _, direction := range directions {
		nextPos := start.Add(direction)
		if !lib.IsPosInBounds(inputMap, nextPos) {
			continue
		}

		if inputMap[nextPos.X][nextPos.Y] == inputMap[start.X][start.Y]+1 {
			ninesReached := findNinesRecursive(inputMap, nextPos)
			for _, nine := range ninesReached {
				nines[nine] = true
			}
		}
	}
	return nines
}

func findNinesRecursive(inputMap [][]rune, pos lib.Point2D) []lib.Point2D {
	if inputMap[pos.X][pos.Y] == '9' {
		return []lib.Point2D{pos}
	}

	nines := make([]lib.Point2D, 0)
	for _, direction := range directions {
		nextPos := pos.Add(direction)
		if !lib.IsPosInBounds(inputMap, nextPos) {
			continue
		}

		if inputMap[nextPos.X][nextPos.Y] == inputMap[pos.X][pos.Y]+1 {
			nines = append(nines, findNinesRecursive(inputMap, nextPos)...)
		}
	}

	return nines
}
