package y2024d12

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

func Part1(matrix [][]rune) int {
	fmt.Printf("%v\n", matrix)
	crawler := NewCrawler(Vec2D{0, 0}, matrix)
	crawler.Crawl()
	return 0
}

type Crawler struct {
	crop              rune
	visitedNodes      Set[Vec2D]
	canBeVisitedNodes Set[Vec2D]
	position          Vec2D
	direcion          Vec2D
	matrix            [][]rune
}

func (cr Crawler) String() string {
	field := make([][]byte, len(cr.matrix[0]))
	for i, row := range cr.matrix {
		field[i] = make([]byte, len(row))
		for j := range field[i] {
			field[i][j] = '.'
		}
	}
	for visited := range cr.visitedNodes.items {
		field[visited.X][visited.Y] = 'O'
	}
	for visited := range cr.canBeVisitedNodes.items {
		field[visited.X][visited.Y] = 'X'
	}
	///
	return ""
}

func NewCrawler(position Vec2D, matrix [][]rune) Crawler {
	return Crawler{
		crop:              matrix[position.X][position.Y],
		visitedNodes:      NewSet[Vec2D](),
		canBeVisitedNodes: NewSet[Vec2D](),
		position:          position,
		direcion:          Vec2D{0, 1},
		matrix:            matrix,
	}
}

func (cr *Crawler) Crawl() {
	cr.CheckNeigbors()
}

func withinBounds(pos Vec2D, matrix [][]rune) bool {
	if pos.X < 0 || pos.X > len(matrix)-1 {
		return false
	}
	if pos.Y < 0 || pos.Y > len(matrix[pos.X])-1 {
		return false
	}
	return true
}

func (cr *Crawler) CheckNeigbors() {
	for _, dir := range [4]Vec2D{
		{0, -1}, // left
		{1, 0},  // down
		{0, 1},  // right
		{-1, 0}, // up
	} {
		if dir.MulScalar(-1) == cr.direcion {
			continue
		}
		nPos := cr.position.Add(dir)
		if cr.CanGo(nPos) {
			cr.canBeVisitedNodes.Add(nPos)
		}
	}
}

func (cr *Crawler) CanGo(pos Vec2D) bool {
	if withinBounds(pos, cr.matrix) && cr.crop == cr.matrix[pos.X][pos.Y] {
		return true
	}
	return false
}

func Part2(matrix [][]rune) int {
	return 0
}

func parseInput(inputPath string) [][]rune {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var matrix [][]rune
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()
		lineRunes := make([]rune, len(line))
		for i, r := range line {
			lineRunes[i] = r
		}
		matrix = append(matrix, lineRunes)
	}
	return matrix
}
