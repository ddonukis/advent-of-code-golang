package router

import (
	"fmt"

	"github.com/ddonukis/advent-of-code-golang/solutions/2018/y2018d01"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d01"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d02"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d03"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d04"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d05"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d06"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d07"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d08"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d09"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d11"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d12"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d13"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d14"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d15"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d16"
	"github.com/ddonukis/advent-of-code-golang/solutions/2024/y2024d19"
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
	case year == 2024 && day == 5:
		y2024d05.Solve(filePath)
	case year == 2024 && day == 6:
		y2024d06.Solve(filePath)
	case year == 2024 && day == 7:
		y2024d07.Solve(filePath)
	case year == 2024 && day == 8:
		y2024d08.Solve(filePath)
	case year == 2024 && day == 9:
		y2024d09.Solve(filePath)
	case year == 2024 && day == 11:
		y2024d11.Solve(filePath)
	case year == 2024 && day == 12:
		y2024d12.Solve(filePath)
	case year == 2024 && day == 13:
		y2024d13.Solve(filePath)
	case year == 2024 && day == 14:
		y2024d14.Solve(filePath)
	case year == 2024 && day == 15:
		y2024d15.Solve(filePath)
	case year == 2024 && day == 16:
		y2024d16.Solve(filePath)
	case year == 2024 && day == 19:
		y2024d19.Solve(filePath)
	default:
		fmt.Printf("No solution found for year %d and day %d.\n", year, day)
	}
}
