package y2024d10

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/ddonukis/advent-of-code-golang/pkg/set"
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

type PartNo int8

const (
	PART_1 PartNo = iota
	PART_2
)

func Part1(grid [][]int8) int {
	startPositions := findStartPositions(grid)

	total := 0
	for _, sp := range startPositions {
		score := exploreTrail(grid, sp, PART_1)
		total += score
	}

	return total
}

func Part2(grid [][]int8) int {
	startPositions := findStartPositions(grid)

	total := 0
	for _, sp := range startPositions {
		score := exploreTrail(grid, sp, PART_2)
		total += score
	}

	return total
}

func findStartPositions(grid [][]int8) []Pos {
	starts := make([]Pos, 0)
	for r, row := range grid {
		for c, height := range row {
			if height == 0 {
				starts = append(starts, Pos{r: int8(r), c: int8(c)})
			}
		}
	}
	return starts
}

// find all 9-height peaks reachable from `start` and return their count
func exploreTrail(grid [][]int8, start Pos, part PartNo) (score int) {
	visited := set.NewSet[Pos]()
	queue := make([]Pos, 1)
	queue[0] = start

	for len(queue) > 0 {
		curPos := queue[0]
		queue = slices.Delete(queue, 0, 1)

		if part == PART_2 || !visited.Contains(curPos) {
			visited.Add(curPos)

			if grid[curPos.r][curPos.c] == 9 {
				score++
			}
		}

		for _, dir := range []Pos{UP, RIGHT, DOWN, LEFT} {
			nextPos := curPos.Add(dir)
			if !visited.Contains(nextPos) && validStep(curPos, nextPos, grid) {
				queue = append(queue, nextPos)
			}
		}
	}
	return
}

func PrintVisited(grid [][]int8, visited set.Set[Pos]) {
	for r, row := range grid {
		for c, i := range row {
			if visited.Contains(Pos{r: int8(r), c: int8(c)}) {
				fmt.Print(i)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func validStep(curPos, nextPos Pos, grid [][]int8) bool {
	if !withinBounds(nextPos, int8(len(grid)), int8(len(grid[0]))) {
		return false
	}
	return grid[curPos.r][curPos.c]+1 == grid[nextPos.r][nextPos.c]
}

func withinBounds(pos Pos, h, w int8) bool {
	return pos.r >= 0 && pos.r < h && pos.c >= 0 && pos.c < w
}

func parseInput(path string) [][]int8 {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimRight(string(content), "\n"), "\n")
	grid := make([][]int8, 0, len(lines))
	for _, line := range lines {
		row := make([]int8, len(line))
		for i, ch := range line {
			n, err := strconv.ParseInt(string(ch), 10, 8)
			if err != nil {
				panic(err)
			}
			row[i] = int8(n)
		}
		grid = append(grid, row)
	}
	return grid
}

// Position
type Pos struct {
	r int8 // row
	c int8 // col
}

func (p Pos) Add(o Pos) Pos {
	return Pos{
		r: p.r + o.r,
		c: p.c + o.c,
	}
}

var RIGHT Pos = Pos{r: 0, c: 1}
var LEFT Pos = Pos{r: 0, c: -1}
var DOWN Pos = Pos{r: 1, c: 0}
var UP Pos = Pos{r: -1, c: 0}
