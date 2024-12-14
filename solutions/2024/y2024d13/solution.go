package y2024d13

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/ddonukis/advent-of-code-golang/pkg/vec"
	"github.com/ddonukis/advent-of-code-golang/solutions/parsing"
)

func Solve(inputPath string) {
	data := parseInput(inputPath)

	t0 := time.Now()
	result1 := Part1(data)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(data)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func parseInput(inputPath string) [][]int {
	fileBytes, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}
	return parsing.ExtractInts(string(fileBytes))
}

type Machine struct {
	A     vec.Vec2D
	B     vec.Vec2D
	Prize vec.Vec2D
}

func (m Machine) TestSolution(a, b int) bool {
	return a*m.A.X+b*m.B.X == m.Prize.X && a*m.A.Y+b*m.B.Y == m.Prize.Y
}

func SolutionCost(a, b int) int {
	return a*3 + b
}

func (m Machine) Optimize() (cost int) {
	minCost := math.MaxInt
	isSolved := false
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			if m.TestSolution(a, b) {
				isSolved = true
				cost := SolutionCost(a, b)
				if minCost > cost {
					minCost = cost
				}
			}
		}
	}
	if isSolved {
		return minCost
	}
	return 0
}

func Part1(data [][]int) int {
	var machines []Machine

	attrs := make([]vec.Vec2D, 0, 3)
	for _, row := range data {
		if len(row) == 2 {
			attrs = append(attrs, vec.Vec2D{X: row[0], Y: row[1]})
		}
		if len(attrs) == 3 {
			machines = append(machines, Machine{
				attrs[0], attrs[1], attrs[2],
			})
			attrs = make([]vec.Vec2D, 0, 3)
		}
	}

	totalCost := 0
	for _, m := range machines {
		// a * Ax + b * Bx = Px
		// a * Ay + b * By = Py
		// a * 3 + b -> min
		totalCost += m.Optimize()
	}

	return totalCost
}

func Part2(data [][]int) int {
	return 0
}
