package lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() string {
	content, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Sprintf("'input' file missing! %v", err))
	}

	return string(content)
}

func ReadInputAsLines() []string {
	return strings.Split(ReadInput(), "\r\n")
}

func ReadMultipleIntValuesPerLine(delimiter string) [][]int {
	lines := ReadInputAsLines()
	values := make([][]int, len(lines))
	for idx, line := range lines {
		valuesInLine := strings.Split(line, delimiter)
		values[idx] = make([]int, len(valuesInLine))
		var err error
		for secIdx, singleValue := range valuesInLine {
			values[idx][secIdx], err = strconv.Atoi(singleValue)
			if err != nil {
				panic(err)
			}
		}
	}

	return values
}
