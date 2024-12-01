package router

import (
	"fmt"

	"github.com/ddonukis/advent-of-code-golang/solutions/2018/y2018d01"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d01"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d02"
)

func RunSolver(year int, day int, filePath string) {
	not_found := false
	switch year {
	case 2018:
		switch day {
		case 1:
			y2018d01.Part1()
		default:
			not_found = true
		}
	case 2024:
		switch day {
		case 1:
			y2024d01.Part1(filePath)
		case 2:
			y2024d02.Part1()
		}
	default:
		not_found = true
	}

	if not_found {
		fmt.Printf("No solution found for year %d and day %d.\n", year, day)
	}
}
