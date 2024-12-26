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
	printGrid(matrix)
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

	printVisited(matrix, visitedToGroupId)

	areaByGroupId := make([]int, groupId)
	sidesByGroupId := make([]int, groupId)

	for pos, groupId := range visitedToGroupId {
		areaByGroupId[groupId] += 1
		sidesByGroupId[groupId] += tileCorners(pos, groupId, visitedToGroupId)
	}

	total := 0
	for i, area := range areaByGroupId {
		p := area * sidesByGroupId[i]
		total += p
		fmt.Printf("group %d: %d * %d = %d\n", i, area, sidesByGroupId[i], p)
	}
	return total
}

// go through each tile in a group (order doesn't matter) and look at all 4 sides:
// if two neigbouring sides (e.g. North and West) are different group -> it's a convex corner
// if two neighbouring sides are same group and the diagonal tile is different group -> it's concave corner
func tileCorners(pos vec.Vec2D, groupId int, tilesByGroupId map[vec.Vec2D]int) int {
	count := 0
	for i := range DIRECTIONS {
		nextPos := pos.Add(DIRECTIONS[i])
		i2 := Mod(i+1, len(DIRECTIONS))
		nextPos2 := pos.Add(DIRECTIONS[i2])

		nextGroupId := getDefault(tilesByGroupId, nextPos, -1)
		nextGroupId2 := getDefault(tilesByGroupId, nextPos2, -1)

		if nextGroupId != groupId && nextGroupId2 != groupId {
			// corner of a convex shape
			count += 1
		} else if nextGroupId == groupId && nextGroupId2 == groupId {
			// corner of a concave shape
			diagonalPos := nextPos.Add(DIRECTIONS[i2])
			if getDefault(tilesByGroupId, diagonalPos, -1) != groupId {
				count += 1
			}
		}
	}
	return count

}

func getDefault[T_K comparable, T_V any](m map[T_K]T_V, key T_K, defVal T_V) T_V {
	val, found := m[key]
	if !found {
		return defVal
	}
	return val
}

func Mod(a, b int) int {
	return (a%b + b) % b
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
