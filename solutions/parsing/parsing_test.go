package parsing

import (
	"fmt"
	"log"
	"slices"
	"testing"
)

func TestExtractInts(t *testing.T) {
	inputs := "10,-20,2\n142,0,\nblah\na=-99;b=(3,2)\n"
	expected := [][]int{
		{10, -20, 2},
		{142, 0},
		{},
		{-99, 3, 2},
	}

	ints := ExtractInts(inputs)

	if len(ints) != len(expected) {
		msg := fmt.Sprintf("Expected slice of len %d, got len of %d\n", len(expected), len(ints))
		fmt.Printf("%v\n", ints)
		log.Fatal(msg)
	}

	for i, row := range ints {
		if !slices.Equal(row, expected[i]) {
			msg := fmt.Sprintf("Row %d: expected %v, got %v \n", i, expected[i], row)
			log.Fatal(msg)
		}
	}
}
