package main

import (
	"aoc2024/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := lib.ReadInputAsLines()
	registerA := lib.MustParseToInt(lines[0][12:])
	registerB := lib.MustParseToInt(lines[1][12:])
	registerC := lib.MustParseToInt(lines[2][12:])

	programCode := strings.Split(lines[4][9:], ",")
	var program []int
	for _, value := range programCode {
		program = append(program, lib.MustParseToInt(value))
	}

	fmt.Println(registerA)
	fmt.Println(registerB)
	fmt.Println(registerC)
	fmt.Println(programCode)
	c := computer{
		registerA:          registerA,
		registerB:          registerB,
		registerC:          registerC,
		instructionPointer: 0,
		program:            program,
		output:             nil,
	}
	c.run()

	var output []string
	for _, out := range c.output {
		output = append(output, strconv.Itoa(out))
	}

	fmt.Printf("Part 1: %s\n", strings.Join(output, ","))
}

type computer struct {
	registerA, registerB, registerC int
	instructionPointer              int
	program                         []int
	output                          []int
}

func (c *computer) run() {
	for ; c.instructionPointer < len(c.program); c.instructionPointer += 2 {
		switch c.program[c.instructionPointer] {
		case 0:
			c.adv(c.program[c.instructionPointer+1])
		case 1:
			c.bxl(c.program[c.instructionPointer+1])
		case 2:
			c.bst(c.program[c.instructionPointer+1])
		case 3:
			c.jnz(c.program[c.instructionPointer+1])
		case 4:
			c.bxc(c.program[c.instructionPointer+1])
		case 5:
			c.out(c.program[c.instructionPointer+1])
		case 6:
			c.bdv(c.program[c.instructionPointer+1])
		case 7:
			c.cdv(c.program[c.instructionPointer+1])
		}
	}
}

func (c *computer) getComboValue(combo int) int {
	if combo <= 3 {
		return combo
	}

	switch combo {
	case 4:
		return c.registerA
	case 5:
		return c.registerB
	case 6:
		return c.registerC
	case 7:
		panic("reserved, do not use")
	}

	panic(fmt.Sprintf("combo %d out of range", combo))
}

func (c *computer) adv(combo int) {
	c.registerA = int(math.Trunc(float64(c.registerA) / math.Pow(2, float64(c.getComboValue(combo)))))
}

func (c *computer) bxl(literal int) {
	c.registerB = c.registerB ^ literal
}

func (c *computer) bst(combo int) {
	c.registerB = c.getComboValue(combo) % 8
}

func (c *computer) jnz(literal int) {
	if c.registerA == 0 {
		return
	}
	// jump to literal and subtract 2 to compensate for loop
	c.instructionPointer = literal - 2
}

func (c *computer) bxc(_ int) {
	c.registerB = c.registerB ^ c.registerC
}

func (c *computer) out(combo int) {
	c.output = append(c.output, c.getComboValue(combo)%8)
}

func (c *computer) bdv(combo int) {
	c.registerB = int(math.Trunc(float64(c.registerA) / math.Pow(2, float64(c.getComboValue(combo)))))
}

func (c *computer) cdv(combo int) {
	c.registerC = int(math.Trunc(float64(c.registerA) / math.Pow(2, float64(c.getComboValue(combo)))))
}
