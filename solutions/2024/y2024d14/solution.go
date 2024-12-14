package y2024d14

import (
	"fmt"
	"os"
	"time"

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

func quadrant(pos Vec2D, H, W int) int {
	halfH, halfW := H/2, W/2
	switch {
	case pos.X < halfH && pos.Y < halfW:
		return 0 // upper left
	case pos.X < halfH && pos.Y > halfW:
		return 1 // upper right
	case pos.X > halfH && pos.Y < halfW:
		return 2 // lower left
	case pos.X > halfH && pos.Y > halfW:
		return 3 // lower right
	}
	return -1
}

func Part1(data [][]int) int {
	const gridHeight, gridWidth = 103, 101
	const numSeconds = 100

	countByQuadrant := [4]int{0, 0, 0, 0}
	for _, inp := range data {
		rob := NewRobot(inp)
		rob.Move(numSeconds, gridHeight, gridWidth)
		q := quadrant(rob.position, gridHeight, gridWidth)
		if q >= 0 && q < 4 {
			countByQuadrant[q]++
		}
	}

	fmt.Printf("%v\n", countByQuadrant)
	total := 1
	for _, n := range countByQuadrant {
		total *= n
	}

	return total
}

func Part2(data [][]int) int {
	return 0
}

type Vec2D struct {
	X int // row
	Y int // col
}

func (v Vec2D) Add(v2 Vec2D) Vec2D {
	return Vec2D{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2D) Wrap(mapHeight int, mapWidth int) Vec2D {
	return Vec2D{Mod(v.X, mapHeight), Mod(v.Y, mapWidth)}
}

func (v Vec2D) MulScalar(n int) Vec2D {
	return Vec2D{v.X * n, v.Y * n}
}

func Mod(a, b int) int {
	return (a%b + b) % b
}

func (v Vec2D) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

type Robot struct {
	position Vec2D
	velocity Vec2D
}

func NewRobot(input []int) Robot {
	if len(input) != 4 {
		msg := fmt.Sprintf("Bad input to NewRobot(): %v\n", input)
		panic(msg)
	}
	return Robot{
		Vec2D{input[1], input[0]},
		Vec2D{input[3], input[2]},
	}
}

func (r *Robot) Move(steps int, mapHeight, mapWidth int) {
	r.position = r.position.Add(r.velocity.MulScalar(steps)).Wrap(mapHeight, mapWidth)

}

func (r Robot) String() string {
	return fmt.Sprintf("Robot{%s, %s}", r.position, r.velocity)
}

func (r Robot) PrintPosition(gridHeight, gridWidth int) {
	var grid [][]byte
	for i := 0; i < gridHeight; i++ {
		row := make([]byte, gridWidth+1)
		for j := range row {
			row[j] = '.'
		}
		row[len(row)-1] = '\n'
		grid = append(grid, row)
	}
	grid[r.position.X][r.position.Y] = 'R'

	for _, row := range grid {
		fmt.Print(string(row))
	}
}
