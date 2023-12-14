package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := os.Args[1]
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fScanner := bufio.NewScanner(f)

	ok := fScanner.Scan()
	if !ok {
		log.Fatal("Empty input file")
	}
	line := fScanner.Text()
	rowIndex := 0
	unmovableCols := make([][]int, len(line))
	movableCols := make([][]int, len(line))
	parseLine(line, rowIndex, unmovableCols, movableCols)
	for fScanner.Scan() {
		rowIndex++
		line = fScanner.Text()
		parseLine(line, rowIndex, unmovableCols, movableCols)
		// fmt.Println(line)
	}

	totalP1 := 0

	// i := 0
	// colLoad := calcColLoad(rowIndex+1, unmovableCols[i], movableCols[i])
	// fmt.Printf("\nCol %d: %d\n\n", i, colLoad)

	for i := 0; i < len(line); i++ {
		colLoad := calcColLoad(rowIndex+1, unmovableCols[i], movableCols[i])
		fmt.Printf("\nCol %d: %d\n\n", i, colLoad)
		totalP1 += colLoad
	}
	fmt.Printf("Total Load (P1): %d\n", totalP1)

}

func parseLine(line string, row int, unmovableCols, movableCols [][]int) {
	for i, r := range line {
		switch r {
		case 'O':
			movableCols[i] = append(movableCols[i], row)
		case '#':
			unmovableCols[i] = append(unmovableCols[i], row)
		default:
			continue
		}
	}
}

func calcLoad(rowIndex, totalRows int) int {
	return totalRows - rowIndex
}

func calcCumLoad(rowIndex, count, totalRows int) int {
	total := 0
	for i := 0; i < count; i++ {
		total += calcLoad(rowIndex+i, totalRows)
	}
	return total
}

func calcColLoad(totalRows int, unmovablePositions, movablePositions []int) int {
	var totalLoad, fromRow, toRow, lastMovableIdx int
	var totalRocks int

	for i := 0; i < len(unmovablePositions)+1; i++ {
		// find how many movable rocks between current unmovable and next unmovable (or end of grid)
		// fmt.Println()
		if i < len(unmovablePositions) {
			toRow = unmovablePositions[i]
		} else {
			toRow = totalRows
		}
		// fmt.Printf("%d -> %d\n", fromRow, toRow)
		if fromRow == toRow {
			continue
		}
		// fmt.Printf("lastMovableIdx: %d\n", lastMovableIdx)
		movablePositions = movablePositions[lastMovableIdx:]
		// fmt.Printf("movablePositions[%d:]\n", lastMovableIdx)
		// fmt.Println(movablePositions)
		totalMovableRocksInRange := 0
		for j := 0; j < len(movablePositions); j++ {
			// fmt.Printf("j: %d\n", j)
			if movablePositions[j] < toRow {
				totalMovableRocksInRange++
			} else {
				lastMovableIdx = j
				// fmt.Printf("lastMovableIdx = %d\n", lastMovableIdx)
				break
			}
		}

		// fmt.Printf("totalMovableRocksInRange: %d\n", totalMovableRocksInRange)
		totalLoad += calcCumLoad(fromRow, totalMovableRocksInRange, totalRows)
		fromRow = toRow + 1
		totalRocks += totalMovableRocksInRange
		// fmt.Printf("load: %d\n", totalLoad)
		if totalMovableRocksInRange == len(movablePositions) {
			break
		}
	}
	// fmt.Printf("movable rocks: %d\n", totalRocks)
	return totalLoad
}
