package router

import (
	"fmt"

	"github.com/ddonukis/advent-of-code-golang/solutions/2018/y2018d01"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d01"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d02"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d03"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d04"
)

func RunSolver(year int, day int, filePath string) {
	switch {
	// 2018
	case year == 2018 && day == 1:
		y2018d01.Solve(filePath)
	// 2024
	case year == 2024 && day == 1:
		y2024d01.Solve(filePath)
	case year == 2024 && day == 2:
		y2024d02.Solve(filePath)
	case year == 2024 && day == 3:
		y2024d03.Solve(filePath)
	case year == 2024 && day == 4:
		y2024d04.Solve(filePath)
	default:
		fmt.Printf("No solution found for year %d and day %d.\n", year, day)
	}
}
