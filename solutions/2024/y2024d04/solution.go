package y2024d04

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Solve(inputPath string) {
	fmt.Printf("AoC 2024 - day 3\n\n")

	charMatrix, err := readCharMatrix(inputPath)
	if err != nil {
		log.Fatalln(err)
	}

	t0 := time.Now()
	result1 := Part1(charMatrix)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(charMatrix)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func readCharMatrix(inputPath string) (charMatrix CharMatrix, err error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	charMatrix = make(CharMatrix, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		charMatrix = append(charMatrix, line)
	}
	return
}

type CharMatrix [][]byte

func (m CharMatrix) String() string {
	capacity := 0
	for _, ln := range m {
		capacity += len(ln)
	}

	elements := make([]string, 0, capacity)
	for _, ln := range m {
		for _, ch := range ln {
			elements = append(elements, fmt.Sprintf("%c, ", ch))
		}
		elements = append(elements, fmt.Sprintln())
	}
	return strings.Join(elements, "")
}

func (m CharMatrix) TurnDiagonalClockwise() CharMatrix {
	rowCount := len(m)

	for row, ln := range m {
		if len(ln) != rowCount {
			log.Fatalf("Matrix must be square, row %d doesn't satisfy the requirment.", row)
		}
	}

	diagonalizedRowCount := (rowCount-1)*2 + 1

	rotatedMatrix := make(CharMatrix, diagonalizedRowCount)

	for row := 0; row < diagonalizedRowCount; row++ {
		var i, j int
		if row < rowCount {
			i = 0
		} else {
			i = row - rowCount + 1
		}
		j = row - i

		for {
			rotatedMatrix[row] = append(rotatedMatrix[row], m[i][j])
			i += 1
			j -= 1
			if i >= rowCount || j < 0 {
				break
			}
		}
	}
	return rotatedMatrix
}

func (m CharMatrix) TurnDiagonalCounterClockwise() CharMatrix {
	rowCount := len(m)

	for row, ln := range m {
		if len(ln) != rowCount {
			log.Fatalf("Matrix must be square, row %d doesn't satisfy the requirment.", row)
		}
	}

	diagonalizedRowCount := (rowCount-1)*2 + 1

	rotatedMatrix := make(CharMatrix, diagonalizedRowCount)

	for row := 0; row < diagonalizedRowCount; row++ {
		var i, j int
		if row < rowCount {
			i = 0
			j = (rowCount - 1) - row
		} else {
			i = row - rowCount + 1
			j = 0
		}

		for {
			rotatedMatrix[row] = append(rotatedMatrix[row], m[i][j])
			i += 1
			j += 1
			if i >= rowCount || j >= rowCount {
				break
			}
		}
	}
	return rotatedMatrix
}

func (m CharMatrix) Transpose() CharMatrix {
	rowCount := len(m)

	for row, ln := range m {
		if len(ln) != rowCount {
			log.Fatalf("Matrix must be square, row %d doesn't satisfy the requirment.", row)
		}
	}

	rotatedMatrix := make(CharMatrix, rowCount)
	for rowIdx := range m {
		rotatedMatrix[rowIdx] = make([]byte, rowCount)
	}

	for rowIdx, row := range m {
		for colIdx, item := range row {
			rotatedMatrix[colIdx][rowIdx] = item
		}
	}
	return rotatedMatrix
}

func Part1(matrix CharMatrix) int {
	count := 0
	for _, ln := range matrix {
		xmas := strings.Count(string(ln), "XMAS")
		samx := strings.Count(string(ln), "SAMX")

		// fmt.Printf("%s : %d\n", string(ln), xmas+samx)

		count = count + xmas + samx
	}

	for _, ln := range matrix.Transpose() {
		xmas := strings.Count(string(ln), "XMAS")
		samx := strings.Count(string(ln), "SAMX")

		// fmt.Printf("%s : %d\n", string(ln), xmas+samx)

		count = count + xmas + samx
	}

	for _, ln := range matrix.TurnDiagonalClockwise() {
		xmas := strings.Count(string(ln), "XMAS")
		samx := strings.Count(string(ln), "SAMX")

		// fmt.Printf("%s : %d\n", string(ln), xmas+samx)

		count = count + xmas + samx
	}

	for _, ln := range matrix.TurnDiagonalCounterClockwise() {
		xmas := strings.Count(string(ln), "XMAS")
		samx := strings.Count(string(ln), "SAMX")

		// fmt.Printf("%s : %d\n", string(ln), xmas+samx)

		count = count + xmas + samx
	}

	return count
}

func Part2(matrix CharMatrix) int {
	return 0
}
