package main

import (
	"aoc2024/lib"
	"errors"
	"fmt"
	"maps"
)

var (
	directions = []lib.Point2D{
		lib.NewPoint2D(-1, 0),
		lib.NewPoint2D(0, 1),
		lib.NewPoint2D(1, 0),
		lib.NewPoint2D(0, -1),
	}
)

var (
	ErrNotInMap = fmt.Errorf("not in map")
)

func main() {
	guardMap := lib.ReadInputAsRuneMap()
	convertDotsToZeros(guardMap)

	currentPos := determineStartPosition(guardMap)
	currentDirection := 0

	tilesVisited := 1
	obstacleOptionSet := make(map[lib.Point2D]any)
	// replace start marker with initial direction of the guard
	guardMap[currentPos.X][currentPos.Y] = 1 << currentDirection
	for {
		// keep moving in the current direction
		var err error
		var nextPos lib.Point2D
		nextPos, currentDirection, err = determineNextTile(guardMap, currentPos, currentDirection)
		if errors.Is(err, ErrNotInMap) {
			break
		}

		if doesTrailLoop(guardMap, currentPos, nextPos, currentDirection) {
			obstacleOptionSet[nextPos] = true
		}

		// store new position, now that we know it's safe
		currentPos = nextPos

		// if we've never been here count this position
		if guardMap[currentPos.X][currentPos.Y] == 0 {
			tilesVisited++
		}

		// save the direction the guard was heading as it passed this point
		guardMap[currentPos.X][currentPos.Y] |= 1 << currentDirection
	}

	fmt.Printf("Day 6 Part 1: %d\n", tilesVisited)
	/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
	obstacleOptions := 0
	for range maps.Keys(obstacleOptionSet) {
		obstacleOptions++
	}
	fmt.Printf("Day 6 Part 2: %d\n", obstacleOptions)

}

func determineNextTile(walkableMap [][]rune, pos lib.Point2D, dir int) (lib.Point2D, int, error) {
	// attempt to take a step in the current direction
	nextPos := pos.Add(directions[dir])
	if !lib.IsPosInBounds(walkableMap, nextPos) {
		// guard left the map, we're done
		return nextPos, -1, ErrNotInMap
	}

	// can't move into walls, change direction and attempt to take a step
	for walkableMap[nextPos.X][nextPos.Y] == '#' {
		dir = (dir + 1) % len(directions)
		// we changed our direction, mark our current position with the new direction
		walkableMap[pos.X][pos.Y] |= 1 << dir

		// move in the new direction
		nextPos = pos.Add(directions[dir])
		if !lib.IsPosInBounds(walkableMap, nextPos) {
			// guard left the map, we're done
			break
		}
	}

	return nextPos, dir, nil
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
		var err error
		var nextPos lib.Point2D
		nextPos, currentDirection, err = determineNextTile(trailMap, currentPos, currentDirection)
		if errors.Is(err, ErrNotInMap) {
			return false
		}

		currentPos = nextPos

		if (trailMap[currentPos.X][currentPos.Y] & (1 << currentDirection)) != 0 {
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
