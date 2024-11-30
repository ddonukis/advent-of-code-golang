package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	squareId, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	width := 0
	prevCycleEndNum := 0
	currentCycleEndNum := 0
	for i := 1; i < squareId*2; i += 2 {
		currentCycleEndNum = i * i
		if i*i >= squareId {
			width = i
			break
		}
		prevCycleEndNum = currentCycleEndNum
	}
	numsOnEdge := currentCycleEndNum - prevCycleEndNum
	fmt.Printf("prev: %d, end: %d, nums: %d\n", prevCycleEndNum, currentCycleEndNum, numsOnEdge)

	offset := squareId - prevCycleEndNum

	fmt.Printf("Offset: %d, corner: %d\n", offset, offset%(width-1))

	distance := abs(offset%(width-1) - (width-1)/2)

	fmt.Printf("from center: %d + %d\n", distance, (width-1)/2)

	fmt.Printf("Part 1: %d\n", distance+(width-1)/2)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
