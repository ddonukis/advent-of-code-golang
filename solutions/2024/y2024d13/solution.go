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
	A vec.Vec2D
	B vec.Vec2D
	P vec.Vec2D
}

func (m Machine) TestSolution(a, b int) bool {
	return a*m.A.X+b*m.B.X == m.P.X && a*m.A.Y+b*m.B.Y == m.P.Y
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

		// a = (Px - b * Bx) / Ax
		// ((Px - b * Bx) / Ax) * Ay + b * By = Py
		// ((Px - b * Bx) / Ax) + b * By / Ay = Py / Ay
		// Px - b * Bx + b * By * Ax / Ay = Py * Ax / Ay
		// - b * Bx + b * By * Ax / Ay = Py * Ax / Ay - Px
		// b * (-1 * Bx + By * Ax / Ay) = Py * Ax / Ay - Px
		// b * (By * Ax / Ay - Bx) = Py * Ax / Ay - Px
		// b = (Py * Ax / Ay - Px) / (By * Ax / Ay - Bx)

		Px := float64(m.P.X)
		Py := float64(m.P.Y)
		Ax := float64(m.A.X)
		Ay := float64(m.A.Y)
		Bx := float64(m.B.X)
		By := float64(m.B.Y)

		// b=(py*ax-px*ay)/(by*ax-bx*ay) a=(px-b*bx)/ax

		b := (Py*Ax - Px*Ay) / (By*Ax - Bx*Ay)
		a := (Px - b*Bx) / Ax

		bIsInt := (b - math.Trunc(b)) < 0.0000001
		aIsInt := (a - math.Trunc(a)) < 0.0000001

		if bIsInt && aIsInt {
			totalCost += int(a)*3 + int(b)
		}

		// totalCost += m.Optimize()
	}

	return totalCost
}

func Part2(data [][]int) int {
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
		m.P.X += 10000000000000
		m.P.Y += 10000000000000
		// a * Ax + b * Bx = Px
		// a * Ay + b * By = Py

		// a = (Px - b * Bx) / Ax

		Px := float64(m.P.X)
		Py := float64(m.P.Y)
		Ax := float64(m.A.X)
		Ay := float64(m.A.Y)
		Bx := float64(m.B.X)
		By := float64(m.B.Y)

		b := (Py*Ax - Px*Ay) / (By*Ax - Bx*Ay)
		a := (Px - b*Bx) / Ax

		bIsInt := (b - math.Trunc(b)) < 0.0000001
		aIsInt := (a - math.Trunc(a)) < 0.0000001

		if bIsInt && aIsInt {
			totalCost += int(a)*3 + int(b)
		}
	}

	return totalCost
}
