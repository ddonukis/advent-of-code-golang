package y2024d14

import "testing"

func TestGridDensity(t *testing.T) {

	// .#.#..
	// .#..#.
	// ..##..
	// ...#..
	// ##....

	// 01000
	// 01000
	// 00120
	// 00010
	// 11000

	grid := [][]bool{
		{false, true, false, true, false, false},
		{false, true, false, false, true, false},
		{false, false, true, true, false, false},
		{false, false, false, true, false, false},
		{true, true, false, false, false, false},
	}

	density := gridDensity(grid)
	if density != 8 {
		t.Fatalf("Expected: 8, got: %d\n", density)
	}
}

func TestGridDensity2(t *testing.T) {

	// .#....
	// .##.#.
	// .###..
	// ..###.
	// #..#..

	// 010000 | 1 | 1
	// 032000 | 5 | 6
	// 024200 | 8 | 14
	// 002410 | 7 | 21
	// 000100 | 1 | 22

	grid := [][]bool{
		{false, true, false, false, false, false},
		{false, true, true, false, true, false},
		{false, true, true, true, false, false},
		{false, false, true, true, true, false},
		{true, false, false, true, false, false},
	}

	density := gridDensity(grid)
	if density != 22 {
		t.Fatalf("Expected: 22, got: %d\n", density)
	}
}
