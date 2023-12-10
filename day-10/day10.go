package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func main() {
	fp, err := getFilePath()
	if err != nil {
		log.Fatalf("Could not get input file path: %v\n", err)
	}

	maze := make(Maze, 0)
	for row := range iterParsedLines(fp, parseRow) {
		maze = append(maze, row)
	}
	fmt.Println(maze)
}

type MazeRow []int8

func (row MazeRow) String() string {
	tiles := make([]rune, len(row))
	for i, tile := range row {
		switch tile {
		case NO_PIPE:
			tiles[i] = ' '
		case VERTICAL:
			tiles[i] = '│'
		case HORIZONTAL:
			tiles[i] = '─'
		case BEND_NORTH_EAST:
			tiles[i] = '└'
		case BEND_NORTH_WEST:
			tiles[i] = '┘'
		case BEND_SOUTH_WEST:
			tiles[i] = '┐'
		case BEND_SOUTH_EAST:
			tiles[i] = '┌'
		case START:
			tiles[i] = 'S'
		}
	}
	return string(tiles)
}

type Maze []MazeRow

func (m Maze) String() string {
	strRows := make([]string, len(m))
	for i, r := range m {
		strRows[i] = r.String()
	}
	return strings.Join(strRows, "\n")
}

func parseRow(line string) MazeRow {
	row := make(MazeRow, 0, len(line))
	for _, tile := range line {
		var tileType int8

		switch tile {
		case '|':
			tileType = VERTICAL
		case '-':
			tileType = HORIZONTAL
		case 'L':
			tileType = BEND_NORTH_EAST
		case 'J':
			tileType = BEND_NORTH_WEST
		case '7':
			tileType = BEND_SOUTH_WEST
		case 'F':
			tileType = BEND_SOUTH_EAST
		case 'S':
			tileType = START
		default: // should only be ground
			tileType = NO_PIPE
		}
		row = append(row, tileType)
	}
	return row
}

const (
	NO_PIPE         int8 = iota // .
	VERTICAL                    // │  |
	HORIZONTAL                  // ─  -
	BEND_NORTH_EAST             // └  L
	BEND_NORTH_WEST             // ┘  J
	BEND_SOUTH_WEST             // ┐  7
	BEND_SOUTH_EAST             // ┌  F
	START                       // S
)
