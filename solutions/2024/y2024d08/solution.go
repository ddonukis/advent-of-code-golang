package y2024d08

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Grid struct {
	anitnodes [][]bool
	antennas  map[rune][]Point2D
	width     int
	height    int
}

func (g *Grid) IsInBounds(p Point2D) bool {
	return p.X >= 0 && p.X < g.height && p.Y >= 0 && p.Y < g.width
}

func (g *Grid) UniqueAntinodes() int {
	total := 0
	for _, row := range g.anitnodes {
		for _, el := range row {
			if el {
				total++
			}
		}
	}
	return total
}

func (g Grid) String() string {
	var elements []string
	elements = append(elements, "antenns:\n")
	for a := range g.antennas {
		el := fmt.Sprintf("  %c: %v\n", a, g.antennas[a])
		elements = append(elements, el)
	}
	elements = append(elements, "antinodes:\n")
	for _, row := range g.anitnodes {
		for _, el := range row {
			symbol := "."
			if el {
				symbol = "#"
			}
			elements = append(elements, symbol)
		}
		elements = append(elements, "\n")
	}

	return strings.Join(elements, "")
}

func (g *Grid) AddAntinode(an Point2D) {
	g.anitnodes[an.X][an.Y] = true
}

func Solve(inputPath string) {
	grid := parseInput(inputPath)

	t0 := time.Now()
	result1 := Part1(grid)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(grid)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

type Point2D struct {
	X, Y int
}

func (p Point2D) DistanceVec(otherP Point2D) Point2D {
	return Point2D{
		otherP.X - p.X,
		otherP.Y - p.Y,
	}
}

func (p Point2D) Add(otherP Point2D) Point2D {
	return Point2D{
		otherP.X + p.X,
		otherP.Y + p.Y,
	}
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func parseInput(inputPath string) Grid {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	antennaLocByFrequency := make(map[rune][]Point2D)
	row := 0
	var antinodes [][]bool
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()
		// fmt.Println(line)
		for col, ch := range line {
			if ch != '.' {
				antennaLocByFrequency[ch] = append(antennaLocByFrequency[ch], Point2D{row, col})
			}
		}
		antinodes = append(antinodes, make([]bool, len(line)))
		row++
	}
	// fmt.Printf("%v\n", antennaLocByFrequency)
	// fmt.Printf("keys: %d\n", len(antennaLocByFrequency))
	return Grid{antinodes, antennaLocByFrequency, len(antinodes), len(antinodes[0])}
}

func Part1(grid Grid) int {
	for freq, ans := range grid.antennas {
		fmt.Printf("%c:\n", freq)
		for curAntIdx := 0; curAntIdx < len(ans); curAntIdx++ {
			curAnt := ans[curAntIdx]
			for compAntIdx := 0; compAntIdx < len(ans); compAntIdx++ {
				if compAntIdx == curAntIdx {
					continue
				}
				dist := curAnt.DistanceVec(ans[compAntIdx])
				fmt.Printf("  %v -> %v: %v\n", curAnt, ans[compAntIdx], dist)
				antinodeLoc := ans[compAntIdx].Add(dist)
				if grid.IsInBounds(antinodeLoc) {
					grid.AddAntinode(antinodeLoc)
				}
			}
		}
	}

	fmt.Println(grid)

	return grid.UniqueAntinodes()
}

func Part2(grid Grid) int {
	for freq, ans := range grid.antennas {
		fmt.Printf("%c:\n", freq)
		for curAntIdx := 0; curAntIdx < len(ans); curAntIdx++ {
			curAnt := ans[curAntIdx]
			if len(ans) > 0 {
				grid.AddAntinode(curAnt)
			}
			for compAntIdx := 0; compAntIdx < len(ans); compAntIdx++ {
				if compAntIdx == curAntIdx {
					continue
				}
				grid.AddAntinode(ans[compAntIdx])

				dist := curAnt.DistanceVec(ans[compAntIdx])
				fmt.Printf("  %v -> %v: %v\n", curAnt, ans[compAntIdx], dist)

				antinodeLoc := ans[compAntIdx].Add(dist)
				for grid.IsInBounds(antinodeLoc) {
					grid.AddAntinode(antinodeLoc)
					antinodeLoc = antinodeLoc.Add(dist)
				}

			}
		}
	}

	fmt.Println(grid)

	return grid.UniqueAntinodes()
}
