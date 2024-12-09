package main

/*
import (
	"aoc2024/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type block struct {
	id     int
	length int
}

func parseInputToBlocks(input string) []block {
	blocks := make([]block, len(input))
	isEmptyBlock := false
	for idx := range input {
		id := idx / 2
		if isEmptyBlock {
			id = -1
		}
		blocks[idx] = block{
			id:     id,
			length: int(input[idx] - '0'),
		}
		isEmptyBlock = !isEmptyBlock
	}
	blocks = append(blocks, block{id: -1, length: 0})
	return blocks
}

func findNextEmptyBlock(blocks []block) int {
	for i := range blocks {
		if blocks[i].id == -1 && blocks[i].length > 0 {
			return i
		}
	}
	panic("no empty blocks remaining!")
}

func moveAsManyBlocksAsPossible(blocks []block, emptyBlockPosition, blocksToMovePosition int) {
	// calculate how many can be moved
	movableBlockCount := int(math.Min(float64(blocks[emptyBlockPosition].length), float64(blocks[blocksToMovePosition].length)))
	// add a block before emptyBlockPosition with the moved blocks and id
	newBlock := block{
		id:     blocks[blocksToMovePosition].id,
		length: movableBlockCount,
	}
	blocks = append(blocks[:emptyBlockPosition], append([]block{newBlock}, blocks[emptyBlockPosition:]...)...)
	// lower emptyBlockPositions+1 length
	blocks[emptyBlockPosition+1].length -= movableBlockCount
	// lower blockToMovePositions+1 length
	blocks[blocksToMovePosition+1].length -= movableBlockCount
	// increase the final blocks length
	blocks[len(blocks)-1].length += movableBlockCount
	//  remove emptyBlockPositions+1if empty
	blocksAddedAndNotRemoved := 1
	if blocks[emptyBlockPosition+1].length == 0 {
		blocks = append(blocks[:emptyBlockPosition+1], blocks[emptyBlockPosition+2:]...)
		blocksAddedAndNotRemoved = 0
	}
	// 	remove blockToMovePosition+1 if empty
	if blocks[blocksToMovePosition+blocksAddedAndNotRemoved].length == 0 {
		blocks = append(blocks[:blocksToMovePosition+blocksAddedAndNotRemoved], blocks[blocksToMovePosition+blocksAddedAndNotRemoved+1:]...)
	}
}

func vizBlocks(blocks []block) {
	for i := range blocks {
		symbol := strconv.Itoa(blocks[i].id)
		if blocks[i].id == -1 {
			symbol = "."
		}
		for range blocks[i].length {
			fmt.Print(symbol)
		}
	}
	fmt.Println()
}

func main() {
	input := lib.ReadInput()
	strings.TrimSpace(input)
	blocks := parseInputToBlocks(input)
	vizBlocks(blocks)

	nextEmptyBlock := findNextEmptyBlock(blocks)
	for nextEmptyBlock != len(blocks) {
		moveAsManyBlocksAsPossible(blocks, nextEmptyBlock, len(blocks)-2)
		vizBlocks(blocks)
	}

}
*/
