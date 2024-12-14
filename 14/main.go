package main

import (
	"aoc2024/lib"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"regexp"
	"strconv"
)

type guardInput struct {
	startPos    lib.Point2D
	movementVec lib.Point2D
}

var (
	mapSize = lib.NewPoint2D(103, 101)
	seconds = 100
)

func main() {
	input := lib.ReadInputAsLines()

	regex := regexp.MustCompile("(\\d+),(\\d+) v=(-?\\d+),(-?\\d+)")
	guards := make([]guardInput, 0, len(input))
	for _, line := range input {
		matches := regex.FindAllStringSubmatch(line, -1)
		//fmt.Printf("%q results in %v\n", line, matches[0][1:])
		guards = append(guards, guardInput{
			startPos:    lib.NewPoint2D(lib.MustParseToInt(matches[0][2]), lib.MustParseToInt(matches[0][1])),
			movementVec: lib.NewPoint2D(lib.MustParseToInt(matches[0][4]), lib.MustParseToInt(matches[0][3])),
		})
	}

	/*
		// Prints the start positions
		for x := 0; x < mapSize.X; x++ {
			for y := 0; y < mapSize.Y; y++ {
				guardCount := 0
				for _, guard := range guards {
					if guard.startPos.X == x && guard.startPos.Y == y {
						guardCount++
					}
				}
				if guardCount == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(strconv.Itoa(guardCount))
				}
			}
			fmt.Print("\n")
		}
	*/

	finalPositions := make([]lib.Point2D, 0, len(guards))
	for _, guard := range guards {
		movement := lib.NewPoint2D(guard.movementVec.X*seconds, guard.movementVec.Y*seconds)
		finalPosition := guard.startPos.Add(movement)
		// modulo to get back into range and then add and modulo once more to turn negative values positive (no effect on positive values)
		finalPosition.X = (finalPosition.X%mapSize.X + mapSize.X) % mapSize.X
		finalPosition.Y = (finalPosition.Y%mapSize.Y + mapSize.Y) % mapSize.Y
		finalPositions = append(finalPositions, finalPosition)
	}

	//printPositions(finalPositions)

	quadrantXSize := (mapSize.X - 1) / 2
	quadrantYSize := (mapSize.Y - 1) / 2

	guardsInEachQuadrant := make([]int, 4)
	for _, guard := range finalPositions {
		if guard.X < quadrantXSize {
			if guard.Y < quadrantYSize {
				guardsInEachQuadrant[0]++
			} else if guard.Y > quadrantYSize {
				guardsInEachQuadrant[1]++
			}
		} else if guard.X > quadrantXSize {
			if guard.Y < quadrantYSize {
				guardsInEachQuadrant[2]++
			} else if guard.Y > quadrantYSize {
				guardsInEachQuadrant[3]++
			}
		}
	}

	/*
		fmt.Println("Quadrant counts:")
		fmt.Printf("%d | %d\n", guardsInEachQuadrant[0], guardsInEachQuadrant[1])
		fmt.Println("~~~~~~~~~")
		fmt.Printf("%d | %d\n", guardsInEachQuadrant[2], guardsInEachQuadrant[3])
	*/
	fmt.Printf("Part 1: %d\n", guardsInEachQuadrant[0]*guardsInEachQuadrant[1]*guardsInEachQuadrant[2]*guardsInEachQuadrant[3])

	/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

	for secondsPassed := 0; secondsPassed < 10000; secondsPassed++ {
		finalPositions = make([]lib.Point2D, 0, len(guards))
		for _, guard := range guards {
			movement := lib.NewPoint2D(guard.movementVec.X*secondsPassed, guard.movementVec.Y*secondsPassed)
			finalPosition := guard.startPos.Add(movement)
			// modulo to get back into range and then add and modulo once more to turn negative values positive (no effect on positive values)
			finalPosition.X = (finalPosition.X%mapSize.X + mapSize.X) % mapSize.X
			finalPosition.Y = (finalPosition.Y%mapSize.Y + mapSize.Y) % mapSize.Y
			finalPositions = append(finalPositions, finalPosition)
		}
		positionsToJpeg(finalPositions, secondsPassed)
	}

}

func printPositions(positions []lib.Point2D) {
	for x := 0; x < mapSize.X; x++ {
		for y := 0; y < mapSize.Y; y++ {
			guardCount := 0
			for _, position := range positions {
				if position.X == x && position.Y == y {
					guardCount++
				}
			}
			if guardCount == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(strconv.Itoa(guardCount))
			}
		}
		fmt.Print("\n")
	}
}

func positionsToJpeg(positions []lib.Point2D, id int) {
	imageBytes := make([]uint8, mapSize.X*mapSize.Y)
	for x := 0; x < mapSize.X; x++ {
		for y := 0; y < mapSize.Y; y++ {
			guardCount := 0
			for _, position := range positions {
				if position.X == x && position.Y == y {
					guardCount++
				}
			}
			if guardCount == 0 {
				imageBytes[x*mapSize.Y+y] = 0
			} else {
				imageBytes[x*mapSize.Y+y] = 255
			}
		}
	}

	img := image.NewGray(image.Rect(0, 0, mapSize.X, mapSize.Y))
	img.Pix = imageBytes

	out, _ := os.Create(fmt.Sprintf("./images/img%d.png", id))
	defer out.Close()

	err := png.Encode(out, img)
	if err != nil {
		log.Println(err)
	}
}
