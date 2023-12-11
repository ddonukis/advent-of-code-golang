package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type GalaxyRow []bool

type Galaxy struct {
	id int
	X  int
	Y  int
}

func (g Galaxy) String() string {
	return fmt.Sprintf("G{id: %d, X: %d, Y: %d}", g.id, g.X, g.Y)
}

func (r GalaxyRow) String() string {
	row := make([]rune, len(r))
	for i, isGalaxy := range r {
		if isGalaxy {
			row[i] = '#'
		} else {
			row[i] = '.'
		}
	}
	return string(row)
}

type GalaxyMap []GalaxyRow

func (m GalaxyMap) String() string {
	strRows := make([]string, len(m))
	for i, r := range m {
		strRows[i] = r.String()
	}
	return strings.Join(strRows, "\n")
}

func main() {
	fp, err := getFilePath()
	if err != nil {
		log.Fatalf("Could not get input file path: %v\n", err)
	}
	galaxyMap := make(GalaxyMap, 0)
	for row := range iterParsedLines(fp, parseRow) {
		galaxyMap = append(galaxyMap, row)
	}

	fmt.Println(galaxyMap)

	galaxies := galaxyCoordinates(galaxyMap)
	sumDitances := 0
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			if g2.id > g1.id {
				// fmt.Printf("%d <> %d\n", g1.id, g2.id)
				sumDitances += manhattanDistance(g1, g2)
			}

		}
	}
	fmt.Printf("Sum distances: %d\n", sumDitances)

}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func manhattanDistance(g1 Galaxy, g2 Galaxy) int {
	return abs(g1.X-g2.X) + abs(g1.Y-g2.Y)
}

func galaxyCoordinates(gm GalaxyMap) []Galaxy {
	emptyRows := make([]int, 0, len(gm))
scanRows:
	for rowIdx, row := range gm {
		for _, isGalaxy := range row {
			if isGalaxy {
				continue scanRows
			}
		}
		emptyRows = append(emptyRows, rowIdx)
	}
	emptyCols := make([]int, 0, len(gm[0]))
scanCols:
	for colIdx := 0; colIdx < len(gm[0]); colIdx++ {
		for rowIdx := 0; rowIdx < len(gm); rowIdx++ {
			if gm[rowIdx][colIdx] {
				continue scanCols
			}
		}
		emptyCols = append(emptyCols, colIdx)
	}
	galaxies := make([]Galaxy, 0, len(gm))
	for r, row := range gm {
		for c, isGalaxy := range row {
			if isGalaxy {
				shiftRowsBy, _ := slices.BinarySearch(emptyRows, r)
				shiftColsBy, _ := slices.BinarySearch(emptyCols, c)
				g := Galaxy{id: len(galaxies), X: r + shiftRowsBy, Y: c + shiftColsBy}
				galaxies = append(galaxies, g)
			}
		}
	}
	return galaxies
}

func parseRow(line string) GalaxyRow {
	row := make(GalaxyRow, 0)
	for _, elem := range line {
		if elem == '#' {
			row = append(row, true)
		} else {
			row = append(row, false)
		}
	}
	return row
}

func getFilePath() (string, error) {
	switch len(os.Args[1:]) {
	case 1:
		return os.Args[1], nil
	case 0:
		return "", errors.New("no arguments passed")
	default:
		return "", errors.New("too many arguments, 1 expected")
	}
}

func iterParsedLines[T any](filepath string, parseFunc func(string) T) <-chan T {
	lines := make(chan T)
	go func() {
		defer close(lines)
		f, err := os.Open(filepath)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
		fs := bufio.NewScanner(f)
		for fs.Scan() {
			lines <- parseFunc(fs.Text())
		}
	}()
	return lines
}
