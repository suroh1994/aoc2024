package main

import (
	"aoc2024/lib"
	"fmt"
	"maps"
)

func main() {
	antennaMap := lib.ReadInputAsRuneMap()
	antennaLocations := findAllAntennas(antennaMap)

	antiNodeLocations := make(map[lib.Point2D]bool)
	// for each frequency
	for _, locations := range antennaLocations {
		// iterate over each antenna location
		for locIdx, antennaA := range locations[:len(locations)-1] {
			// iterate over all remaining antenna locations
			for _, antennaB := range locations[locIdx+1:] {
				// calculate all antinodes for the antenna pair
				antiNodeLocationsForPair := calculateAntinodeLocations(antennaA, antennaB)
				// only not antinodes which are in bounds
				for _, antiNodeLoc := range antiNodeLocationsForPair {
					if lib.IsPosInBounds(antennaMap, antiNodeLoc) {
						antiNodeLocations[antiNodeLoc] = true
					}
				}
			}
		}
	}

	uniqueAntinodeLocationCount := 0
	for range maps.Keys(antiNodeLocations) {
		uniqueAntinodeLocationCount++
	}

	fmt.Printf("Part 1: %d\n", uniqueAntinodeLocationCount)
}

func findAllAntennas(antennaMap [][]rune) map[rune][]lib.Point2D {
	antennaLocations := make(map[rune][]lib.Point2D)
	for x := range antennaMap {
		for y := range antennaMap[x] {
			runeAtXY := antennaMap[x][y]
			if (runeAtXY >= '0' && runeAtXY <= '9') ||
				(runeAtXY >= 'A' && runeAtXY <= 'Z') ||
				(runeAtXY >= 'a' && runeAtXY <= 'z') {
				antennaLocations[runeAtXY] = append(antennaLocations[runeAtXY], lib.NewPoint2D(x, y))
			}
		}
	}
	return antennaLocations
}

func calculateAntinodeLocations(antennaA lib.Point2D, antennaB lib.Point2D) []lib.Point2D {
	distX := antennaA.X - antennaB.X
	distY := antennaA.Y - antennaB.Y
	return []lib.Point2D{
		lib.NewPoint2D(antennaA.X+distX, antennaA.Y+distY),
		lib.NewPoint2D(antennaB.X-distX, antennaB.Y-distY),
	}
}
