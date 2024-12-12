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

	field := make([][]RegionNode, len(matrix))
	for rowIdx, row := range matrix {
		field[rowIdx] = make([]RegionNode, len(row))
		// fmt.Printf("%v\n", field)
		fmt.Printf("rowIdx: %d, len(row): %d, len(field): %d len(field[rowIdx]): %d\n", rowIdx, len(row), len(field), len(field[rowIdx]))
		for colIdx, element := range row {
			field[rowIdx][colIdx] = RegionNode{
				crop:      element,
				postition: Vec2D{rowIdx, colIdx},
				regionId:  -1,
			}
		}
	}
	fmt.Println("Created nodes.")
	for i := range field {
		for j := range field[i] {
			if j-1 >= 0 {
				field[i][j].LinkLeft(&field[i][j-1])
			}
			if j+1 < len(field[i]) {
				field[i][j].LinkRight(&field[i][j+1])
			}
			if i-1 >= 0 {
				field[i][j].LinkUp(&field[i-1][j])
			}
			if i+1 < len(field) {
				field[i][j].LinkBottom(&field[i+1][j])
			}
		}
	}
	fmt.Println("Linked nodes.")
	curRegion := 0
	for _, row := range field {
		for _, node := range row {
			if node.regionId == -1 {
				node.PushRegionId(curRegion)
				curRegion++
			}
		}
	}
	fmt.Println("Assigned regions.")
	for _, row := range field {
		for _, node := range row {
			fmt.Printf("%d", node.regionId)
		}
		fmt.Println()
	}

	return 0
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
