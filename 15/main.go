package main

import (
	"aoc2024/lib"
	"fmt"
)

func main() {
	input := lib.ReadInputAsLines()
	divisorLine := 0
	for ; divisorLine < len(input); divisorLine++ {
		if input[divisorLine] == "" {
			break
		}
	}

	warehouseMap := lib.LinesToRuneMap(input[:divisorLine])
	instructionLines := input[divisorLine+1:]
	instructions := ""
	for _, line := range instructionLines {
		instructions += line
	}

	startPos := findStartPos(warehouseMap)

	moveMap := map[rune]lib.Point2D{
		'<': lib.LEFT,
		'>': lib.RIGHT,
		'^': lib.UP,
		'v': lib.DOWN,
	}
	for _, instruction := range instructions {
		move := moveMap[instruction]
		startPos = moveRecursive(warehouseMap, startPos, move)
		//printMap(warehouseMap)
	}

	sumOfCoords := 0
	for x := 0; x < len(warehouseMap); x++ {
		for y := 0; y < len(warehouseMap[x]); y++ {
			if warehouseMap[x][y] == 'O' {
				sumOfCoords += 100*x + y
			}
		}
	}
	fmt.Printf("Part 1: %d\n", sumOfCoords)
}

func findStartPos(warehouseMap [][]rune) lib.Point2D {
	for x := 0; x < len(warehouseMap); x++ {
		for y := 0; y < len(warehouseMap[x]); y++ {
			if warehouseMap[x][y] == '@' {
				return lib.Point2D{X: x, Y: y}
			}
		}
	}
	panic("no guard found")
}

func moveRecursive(warehouseMap [][]rune, position, direction lib.Point2D) lib.Point2D {
	nextPosition := position.Add(direction)
	if warehouseMap[nextPosition.X][nextPosition.Y] == 'O' {
		moveRecursive(warehouseMap, nextPosition, direction)
	}

	if warehouseMap[nextPosition.X][nextPosition.Y] == '[' || warehouseMap[nextPosition.X][nextPosition.Y] == ']' {

		moveRecursive(warehouseMap, nextPosition, direction)
	}

	if warehouseMap[nextPosition.X][nextPosition.Y] == '.' {
		warehouseMap[nextPosition.X][nextPosition.Y] = warehouseMap[position.X][position.Y]
		warehouseMap[position.X][position.Y] = '.'
		return nextPosition
	}

	return position
}

func printMap(warehouseMap [][]rune) {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for x := 0; x < len(warehouseMap); x++ {
		for y := 0; y < len(warehouseMap[x]); y++ {
			fmt.Printf("%c", warehouseMap[x][y])
		}
		fmt.Println()
	}
}

func scaleMap(warehouse [][]rune) [][]rune {
	newWarehouse := make([][]rune, len(warehouse)*2)
	for x := 0; x < len(warehouse); x++ {
		newWarehouse[x] = make([]rune, len(warehouse[x])*2)
		for y := 0; y < len(warehouse[x]); y++ {
			switch warehouse[x][y] {
			case '#':
				newWarehouse[x][2*y] = '#'
				newWarehouse[x][2*y+1] = '#'
			case 'O':
				newWarehouse[x][2*y] = '['
				newWarehouse[x][2*y+1] = ']'
			case '.':
				newWarehouse[x][2*y] = '.'
				newWarehouse[x][2*y+1] = '.'
			case '@':
				newWarehouse[x][2*y] = '@'
				newWarehouse[x][2*y+1] = '.'
			}
		}
	}
	return newWarehouse
}
