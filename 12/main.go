package main

import (
	"aoc2024/lib"
	"fmt"
)

type plot struct {
	crop        rune
	fencedSides int
}

type plotRegion struct {
	area          int
	perimeter     int
	straightEdges int
}

func main() {
	farmMap := lib.ReadInputAsRuneMap()
	plotMap := plotFarmMap(farmMap)

	regions := determineRegions(plotMap)
	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.area * region.perimeter
	}
	fmt.Printf("Part 1: %d\n", totalPrice)
}

func plotFarmMap(farmMap [][]rune) [][]plot {
	plotMap := make([][]plot, len(farmMap))
	for x := 0; x < len(farmMap); x++ {
		plotMap[x] = make([]plot, len(farmMap[x]))
		for y := 0; y < len(farmMap[x]); y++ {
			plotMap[x][y] = plot{
				crop:        farmMap[x][y],
				fencedSides: sameNeighbours(farmMap, lib.NewPoint2D(x, y)),
			}
		}
	}
	return plotMap
}

var (
	neighbouringDirections = []lib.Point2D{
		lib.UP,
		lib.RIGHT,
		lib.DOWN,
		lib.LEFT,
	}
)

func sameNeighbours(farmMap [][]rune, plotPos lib.Point2D) int {
	fencedSides := 4
	cropType := farmMap[plotPos.X][plotPos.Y]
	for _, direction := range neighbouringDirections {
		neighbour := plotPos.Add(direction)
		if lib.IsPosInBounds(farmMap, neighbour) && farmMap[neighbour.X][neighbour.Y] == cropType {
			fencedSides--
		}
	}

	return fencedSides
}

func determineRegions(plotMap [][]plot) []plotRegion {
	plotsAlreadyMapped := make(map[lib.Point2D]bool)
	regions := make([]plotRegion, 0)
	for x := 0; x < len(plotMap); x++ {
		for y := 0; y < len(plotMap[x]); y++ {
			currentPos := lib.Point2D{X: x, Y: y}
			// only add to a region once!
			if _, exists := plotsAlreadyMapped[currentPos]; exists {
				continue
			}

			plotsInRegion := findPointsInRegion(plotMap, &plotsAlreadyMapped, currentPos)
			perimeter := 0
			for _, plots := range plotsInRegion {
				perimeter += plotMap[plots.X][plots.Y].fencedSides
			}
			regions = append(regions, plotRegion{
				area:          len(plotsInRegion),
				perimeter:     perimeter,
				straightEdges: countEdgesInRegion(plotsInRegion),
			})
		}
	}
	return regions
}

func findPointsInRegion(plotMap [][]plot, visitedMap *map[lib.Point2D]bool, currentPos lib.Point2D) []lib.Point2D {
	if (*visitedMap)[currentPos] {
		return nil
	}

	(*visitedMap)[currentPos] = true

	pointsToBeAddedToRegion := []lib.Point2D{currentPos}
	for _, direction := range neighbouringDirections {
		neighbour := currentPos.Add(direction)
		if lib.IsPosInBounds(plotMap, neighbour) &&
			plotMap[currentPos.X][currentPos.Y].crop == plotMap[neighbour.X][neighbour.Y].crop {
			// add neighbour (and neighbour's neighbour) to region
			pointsToBeAddedToRegion = append(pointsToBeAddedToRegion, findPointsInRegion(plotMap, visitedMap, neighbour)...)
		}
	}

	return pointsToBeAddedToRegion
}

func countEdgesInRegion(plotMap [][]plot, plotsInRegion []lib.Point2D) int {
	// get first point in region
	currentPos := plotsInRegion[0]

	// plots with a single tile are easy
	if plotMap[currentPos.X][currentPos.Y].fencedSides == 4 {
		return 4
	}

	// try to find a direction (in clockwise order) you can walk to another point in region
	directionIdx := 0
	currentDir := neighbouringDirections[directionIdx]
	nextPos := currentPos.Add(currentDir)
	for plotMap[currentPos.X][currentPos.Y].crop != plotMap[nextPos.X][nextPos.Y].crop {
		directionIdx++
		currentDir = neighbouringDirections[directionIdx]
		nextPos = currentPos.Add(currentDir)
	}

	// remember the start direction so we know whether we have to add more turns
	startDirection := directionIdx

	currentPos = nextPos
	edgeCount := 0
	// keep walking, until...
	for {
		//  we reached the start
		//	 stop walking
		//   add extra edges if necessary
		//   return edgeCount
		//	next point not in region
		// 	 change direction and add 1 to edge count
		//  a point to your left is walkable
		//   change direction and add 1 to edge count

	}

}
