package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Rock struct {
	movable  bool
	position int
}

type Column []Rock

type Board []Column

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
	board := make(Board, len(line))
	parseLine(line, rowIndex, board)

	for fScanner.Scan() {
		rowIndex++
		line = fScanner.Text()
		parseLine(line, rowIndex, board)
	}

	totalP1 := 0
	for i := 0; i < len(line); i++ {
		colLoad := calcColLoad(rowIndex+1, board[i])
		fmt.Printf("\nCol %d: %d\n\n", i, colLoad)
		totalP1 += colLoad
	}
	fmt.Printf("Total Load (P1): %d\n", totalP1)

}

func parseLine(line string, row int, board Board) {
	for i, r := range line {
		switch r {
		case 'O':
			board[i] = append(board[i], Rock{true, row})
		case '#':
			board[i] = append(board[i], Rock{false, row})
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

func calcColLoad(totalRows int, col Column) int {
	totalLoad := 0
	rangeStart := 0
	movableInRange := 0
	for _, rock := range col {
		if rock.movable {
			movableInRange++
		} else {
			totalLoad += calcCumLoad(rangeStart, movableInRange, totalRows)
			rangeStart = rock.position + 1
			movableInRange = 0
		}
	}
	if movableInRange > 0 {
		totalLoad += calcCumLoad(rangeStart, movableInRange, totalRows)
	}
	return totalLoad
}

func calcColLoad(totalRows int, unmovablePositions, movablePositions []int) int {
	var totalLoad, fromRow, toRow, lastMovableIdx int
	var totalRocks int

	for i := 0; i < len(unmovablePositions)+1; i++ {
		// find how many movable rocks between current unmovable and next unmovable (or end of grid)
		if i < len(unmovablePositions) {
			toRow = unmovablePositions[i]
		} else {
			toRow = totalRows
		}
		if fromRow == toRow {
			continue
		}
		movablePositions = movablePositions[lastMovableIdx:]
		totalMovableRocksInRange := 0
		for j := 0; j < len(movablePositions); j++ {
			if movablePositions[j] < toRow {
				totalMovableRocksInRange++
			} else {
				lastMovableIdx = j
				break
			}
		}
		totalLoad += calcCumLoad(fromRow, totalMovableRocksInRange, totalRows)
		fromRow = toRow + 1
		totalRocks += totalMovableRocksInRange
		if totalMovableRocksInRange == len(movablePositions) {
			break
		}
	}
	return totalLoad
}
