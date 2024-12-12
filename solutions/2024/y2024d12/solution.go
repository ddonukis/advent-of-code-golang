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

	crawlMap := make([][]int, len(matrix))
	for rowIdx := range matrix {
		crawlMap[rowIdx] = make([]int, len(matrix[rowIdx]))
		for colIdx := range crawlMap[rowIdx] {
			crawlMap[rowIdx][colIdx] = -1
		}
	}

	crawlerId := 0
	x, y := 0
	var nextX, nextY int
	for {
		// explore all connected tiles
		if crawlMap[x][y] == -1 {
			crawlMap[x][y] = crawlerId
		}

	}
}

type Vec2D struct {
	X, Y int
}

func (v Vec2D) Add(otherV Vec2D) Vec2D {
	return Vec2D{v.X + otherV.X, v.Y + otherV.Y}
}

func ExploreNodes(start Vec2D, crawlerId CrawlerId, matrix [][]rune, crawlMap [][]int) {
	crawler := NewCrawler(crawlerId, start, matrix)
}

type CrawlerId int

type CrawlState struct {
	crawlerID CrawlerId
	availableDirs
}


type Crawler struct {
	id       CrawlerId
	crop     rune
	position Vec2D
	matrix   [][]rune
	crawlMap [][]CrawlerId
}

func NewCrawler(id CrawlerId, position Vec2D, matrix [][]rune) Crawler {
	return Crawler{
		id:       id,
		crop:     matrix[position.X][position.Y],
		position: position,
		matrix:   matrix,
	}
}

func (cr *Crawler) TryMove(dir Vec2D) bool {
	nextPos := cr.position.Add(dir)
	if !(nextPos.X >= 0 && nextPos.X < len(cr.matrix) && nextPos.Y >= 0 && nextPos.Y < len(cr.matrix[nextPos.X])) {
		return false
	}

	if cr.crop != cr.matrix[nextPos.X][nextPos.Y] {
		return false
	}

	if

	cr.position = nextPos
	return true
}

func (crawler *Crawler) Crawl() {
	for {

	}
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
