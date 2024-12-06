package main

import (
	"aoc2024/lib"
	"fmt"
	"maps"
	"os"
)

var (
	debug      = true
	directions = []lib.Point2D{
		lib.NewPoint2D(-1, 0),
		lib.NewPoint2D(0, 1),
		lib.NewPoint2D(1, 0),
		lib.NewPoint2D(0, -1),
	}
)

func main() {
	guardMap := lib.ReadInputAsRuneMap()
	convertDotsToZeros(guardMap)

	startPos := determineStartPosition(guardMap)
	currentPos := startPos
	currentDirection := 0

	tilesVisited := 1
	obstacleOptionSet := make(map[lib.Point2D]any)
	// replace start marker with initial direction of the guard
	guardMap[currentPos.X][currentPos.Y] = 1 << currentDirection
	for {
		// keep moving in the current direction
		nextPos := currentPos.Add(directions[currentDirection])
		if !lib.IsPosInBounds(guardMap, nextPos) {
			// guard left the map, we're done
			break
		}

		// can't move into walls, change direction and take another step
		if guardMap[nextPos.X][nextPos.Y] == '#' {
			currentDirection = (currentDirection + 1) % len(directions)
			// we changed our direction, mark our current position with the new direction
			guardMap[currentPos.X][currentPos.Y] |= 1 << currentDirection

			// move in the new direction
			nextPos = currentPos.Add(directions[currentDirection])
			if !lib.IsPosInBounds(guardMap, nextPos) {
				// guard left the map, we're done
				break
			}
		}

		// store new position, now that we know it's safe
		currentPos = nextPos

		// if we've never been here count this position
		if guardMap[currentPos.X][currentPos.Y] == 0 {
			tilesVisited++
		}

		// save the direction the guard was heading as it passed this point
		guardMap[currentPos.X][currentPos.Y] |= 1 << currentDirection
		visualizeMap(guardMap, currentPos)

		// check whether an obstacle in front of us would result in a loop
		obstaclePos := currentPos.Add(directions[currentDirection])
		if doesTrailLoop(guardMap, currentPos, obstaclePos, currentDirection) {
			obstacleOptionSet[obstaclePos] = true
		}
	}

	fmt.Printf("Day 6 Part 1: %d\n", tilesVisited)
	/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
	obstacleOptions := 0
	for range maps.Keys(obstacleOptionSet) {
		obstacleOptions++
	}
	fmt.Printf("Day 6 Part 2: %d\n", obstacleOptions)

}

// doesTrailLoop takes the original map, makes a copy of it and places an obstacle (#) at the given location.
// We then follow the same rules and see whether the guard walks back into a path it has taken before or leaves the map
func doesTrailLoop(guardMap [][]rune, startPos, obstaclePos lib.Point2D, startingDirection int) bool {
	// obstacle would be out of bounds, no need to check
	if !lib.IsPosInBounds(guardMap, obstaclePos) {
		return false
	}

	// we cannot place an obstacle at:
	// - positions with obstacles already there
	// - positions we have already passed through, this would change our trajectory
	if guardMap[obstaclePos.X][obstaclePos.Y] != 0 {
		return false
	}

	// create a map copy so we don't ruin the input map
	trailMap := make([][]rune, len(guardMap))
	for i := range guardMap {
		trailMap[i] = make([]rune, len(guardMap[i]))
		copy(trailMap[i], guardMap[i])

		// add the obstacle
		if i == obstaclePos.X {
			trailMap[i][obstaclePos.Y] = '#'
		}
	}

	currentPos := startPos
	currentDirection := startingDirection
	for {
		nextPos := currentPos.Add(directions[currentDirection])
		if !lib.IsPosInBounds(trailMap, nextPos) {
			return false
		}

		if trailMap[nextPos.X][nextPos.Y] == '#' {
			currentDirection = (currentDirection + 1) % len(directions)
			// we changed our direction, mark our current position with the new direction
			trailMap[currentPos.X][currentPos.Y] |= 1 << currentDirection
			// calculate the next position with the updated direction
			nextPos = currentPos.Add(directions[currentDirection])

			if !lib.IsPosInBounds(trailMap, nextPos) {
				return false
			}
		}

		currentPos = nextPos

		if (trailMap[currentPos.X][currentPos.Y] & (1 << currentDirection)) != 0 {
			visualizeMap(trailMap, startPos)
			return true
		}

		trailMap[currentPos.X][currentPos.Y] |= 1 << currentDirection
	}
}

func convertDotsToZeros(guardMap [][]rune) {
	for i := 0; i < len(guardMap); i++ {
		for j := 0; j < len(guardMap[i]); j++ {
			if guardMap[i][j] == '.' {
				guardMap[i][j] = 0
			}
		}
	}
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

func visualizeMap(mapToVisualize [][]rune, currentPos lib.Point2D) {
	if !debug {
		return
	}
	visualization := ""
	for x, row := range mapToVisualize {
		for y, cell := range row {
			if currentPos.X == x && currentPos.Y == y {
				visualization += "O"
				continue
			}

			if cell == '#' {
				visualization += "#"
				continue
			}

			hasPassedVertically := (cell & 5) != 0
			hasPassedHorizontally := (cell & 10) != 0

			symbol := "."
			switch true {
			case hasPassedVertically && hasPassedHorizontally:
				symbol = "+"
			case hasPassedVertically:
				symbol = "|"
			case hasPassedHorizontally:
				symbol = "-"
			}
			visualization += symbol
		}
		visualization += "\n"
	}

	err := os.WriteFile("myMap", []byte(visualization), 0666)
	if err != nil {
		panic(err)
	}
}
