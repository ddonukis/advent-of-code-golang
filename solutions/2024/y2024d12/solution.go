package y2024d12

import (
	"fmt"
	"slices"
	"time"

	"github.com/ddonukis/advent-of-code-golang/pkg/vec"
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
	// printGrid(matrix)

	visitedToGroupId := make(map[vec.Vec2D]int)

	groupId := 0
	for r, row := range matrix {
		for c := range row {
			startPos := vec.Vec2D{X: r, Y: c}
			_, exists := visitedToGroupId[startPos]
			if exists {
				continue
			}
			expandTileGroup(startPos, groupId, matrix, visitedToGroupId)
			groupId++
		}
	}

	// printVisited(matrix, visitedToGroupId)

	areaByGroupId := make([]int, groupId)
	perimterByGroupId := make([]int, groupId)

	for pos, groupId := range visitedToGroupId {
		areaByGroupId[groupId] += 1
		perimterByGroupId[groupId] += tilePerimeter(pos, groupId, visitedToGroupId)
	}

	total := 0
	for i, area := range areaByGroupId {
		// fmt.Printf("group: %d, area: %d, perimiter: %d\n", i, area, perimterByGroupId[i])
		total += area * perimterByGroupId[i]
	}
	return total
}

func Part2(matrix [][]rune) int {
	// fmt.Printf("%v\n", matrix)
	return 0
}

func tilePerimeter(pos vec.Vec2D, groupId int, tilesByGroupId map[vec.Vec2D]int) int {
	p := 0
	for _, dir := range DIRECTIONS {
		nextPos := pos.Add(dir)

		nextGroupId, found := tilesByGroupId[nextPos]
		if !found || (groupId != nextGroupId) {
			p++
		}
	}
	return p
}

func expandTileGroup(
	startPos vec.Vec2D,
	groupId int,
	grid [][]rune,
	visited map[vec.Vec2D]int,
) {
	queue := []vec.Vec2D{startPos}

	for len(queue) > 0 {
		curPos := queue[0]
		queue = slices.Delete(queue, 0, 1)
		_, alreadyVisisted := visited[curPos]
		if alreadyVisisted {
			continue
		}
		visited[curPos] = groupId

		for _, dir := range DIRECTIONS {
			nextPos := curPos.Add(dir)

			_, isVisited := visited[nextPos]

			if isPossibleMove(nextPos, curPos, grid) && !isVisited {
				queue = append(queue, nextPos)
			}
		}
	}
}

var DIRECTIONS = [4]vec.Vec2D{
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
}

func isPossibleMove(nextPos, curPos vec.Vec2D, grid [][]rune) bool {
	if nextPos.X < 0 || nextPos.X >= len(grid) || nextPos.Y < 0 || nextPos.Y >= len(grid[0]) {
		return false
	}
	return grid[curPos.X][curPos.Y] == grid[nextPos.X][nextPos.Y]
}
