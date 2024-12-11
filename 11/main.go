package main

import (
	"aoc2024/lib"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := lib.ReadInput()
	numberStrings := strings.Split(input, " ")

	numbers := make([]int, len(numberStrings), 10000000) // let's start higher...
	for idx := range numberStrings {
		numbers[idx] = lib.MustParseToInt(numberStrings[idx])
	}

	//printStones(numbers)

	for i := 0; i < 25; i++ {
		fmt.Printf("Iteration: %d\n", i+1)
		ApplyRules(&numbers)
		//printStones(numbers)
	}

	fmt.Printf("Part 1: %d\n", len(numbers))
}

var (
	cache = map[int][]int{
		0: {1},
	}
)

func ApplyRules(numbers *[]int) {
	for i := 0; i < len(*numbers); i++ {
		switch true {
		case (*numbers)[i] == 0:
			(*numbers)[i] = 1
		case lib.DigitsInNum((*numbers)[i])%2 == 0:
			left := (*numbers)[i] / int(math.Pow(10, float64(lib.DigitsInNum((*numbers)[i])/2)))
			right := (*numbers)[i] % int(math.Pow(10, float64(lib.DigitsInNum((*numbers)[i])/2)))
			(*numbers)[i] = left
			*numbers = append((*numbers)[:i+1], append([]int{right}, (*numbers)[i+1:]...)...)
			i++
		default:
			(*numbers)[i] = (*numbers)[i] * 2024
		}
	}
}

func printStones(numbers []int) {
	for _, number := range numbers {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}
