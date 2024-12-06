package main

import (
	"aoc2024/lib"
	"fmt"
)

var (
	currentDirection = 0
	directions       = []lib.Point2D{
		lib.NewPoint2D(-1, 0),
		lib.NewPoint2D(0, 1),
		lib.NewPoint2D(1, 0),
		lib.NewPoint2D(0, -1),
	}
)

func main() {
	guardMap := lib.ReadInputAsRuneMap()

	currentPos := determineStartPosition(guardMap)

	tilesVisited := 1
	guardMap[currentPos.X][currentPos.Y] = '+'
	for {
		nextPos := currentPos.Add(directions[currentDirection])
		if !lib.IsPosInBounds(guardMap, nextPos) {
			break
		}

		// can't move into walls, change direction and take another step
		if guardMap[nextPos.X][nextPos.Y] == '#' {
			currentDirection = (currentDirection + 1) % len(directions)
			nextPos = currentPos.Add(directions[currentDirection])

			if !lib.IsPosInBounds(guardMap, nextPos) {
				break
			}
		}

		currentPos = nextPos

		if guardMap[currentPos.X][currentPos.Y] == '.' {
			tilesVisited++
			guardMap[currentPos.X][currentPos.Y] = '*'
		}
	}

	fmt.Printf("Day 6 Part 1: %d\n", tilesVisited)
}

func determineStartPosition(guardMap [][]rune) lib.Point2D {
	for x, row := range guardMap {
		for y, cell := range row {
			if cell == '^' {
				return lib.NewPoint2D(x, y)
			}
		}
	}
	panic("The guard was a lie!")
}
