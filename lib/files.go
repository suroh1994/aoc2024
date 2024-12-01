package lib

import (
	"fmt"
	"os"
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
