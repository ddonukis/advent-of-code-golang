package y2024d12

import (
	"bufio"
	"os"
)

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
