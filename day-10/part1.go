package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
)

func part1(maze Maze) {
	maze = padMaze(maze)
	mazeWidth := len(maze[0])
	mazeHeight := len(maze)

	fmt.Printf("Maze dimensions (WxH): %d x %d\n", mazeWidth, mazeHeight)
	fmt.Printf("Total tiles: %d\n", mazeWidth*mazeHeight)
	start, err := locateStart(maze)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Start location (X, Y): %v\n", start)
	fmt.Println()
	startTile := Tile{pipeType: START, coordinates: start}

	current := startTile
	previous := startTile
	var next Tile
	var exists bool = true

	loopLength := 0
	for i := 0; i < mazeHeight*mazeWidth+1; i++ {
		fmt.Printf("%d: %v\n", i, current)
		exists, next = nextTile(maze, current, previous)
		if !exists {
			fmt.Println("Reached dead end!")
			break
		}
		if next == startTile {
			fmt.Println("Back at start!")
			loopLength = i + 2
			break
		}
		previous = current
		current = next
	}
	fmt.Printf("Loop length: %d\n", loopLength)
	fmt.Printf("Farthest tile at %d steps\n", loopLength/2)

}

type Coordinates struct {
	X int // horizontal
	Y int // vertical
}

type Tile struct {
	coordinates Coordinates
	pipeType    int8
}

func locateStart(maze Maze) (Coordinates, error) {
	for y := range maze {
		for x := range maze[y] {
			if maze[y][x] == START {
				return Coordinates{x, y}, nil
			}
		}
	}
	return Coordinates{}, errors.New("could not find a start point")
}

func padMaze(maze Maze) Maze {
	for i, row := range maze {
		updatedRow := make(MazeRow, 1, len(row)+2)
		updatedRow[0] = NO_PIPE
		updatedRow = append(updatedRow, row...)
		updatedRow = append(updatedRow, NO_PIPE)
		maze[i] = updatedRow
	}

	noPipeRow := make(MazeRow, len(maze[0]))
	for i := range noPipeRow {
		noPipeRow[i] = NO_PIPE
	}

	maze = slices.Insert(maze, 0, noPipeRow)

	return append(maze, slices.Clone(noPipeRow))
}

type Direction int8

const (
	SOUTH Direction = iota
	EAST
	NORTH
	WEST
)

type Move struct {
	direction Direction
	offset    Coordinates
}

var MOVE_PRIORITY [4]Move = [4]Move{
	{SOUTH, Coordinates{X: 0, Y: 1}},
	{EAST, Coordinates{X: 1, Y: 0}},
	{NORTH, Coordinates{X: 0, Y: -1}},
	{WEST, Coordinates{X: -1, Y: 0}},
}

func nextTile(maze Maze, currentTile, previousTile Tile) (exists bool, nextTile Tile) {
selectMove:
	for _, move := range MOVE_PRIORITY {
		nextTileCoord := Coordinates{
			X: currentTile.coordinates.X + move.offset.X,
			Y: currentTile.coordinates.Y + move.offset.Y,
		}
		nextTileType := maze[nextTileCoord.Y][nextTileCoord.X]
		nextTile := Tile{pipeType: nextTileType, coordinates: nextTileCoord}

		if nextTile == previousTile {
			continue selectMove
		}
		switch move.direction {
		case SOUTH:
			if t := currentTile.pipeType; t != VERTICAL && t != BEND_SOUTH_EAST && t != BEND_SOUTH_WEST && t != START {
				continue selectMove
			}

			switch nextTileType {
			case VERTICAL, BEND_NORTH_EAST, BEND_NORTH_WEST, START:
				return true, nextTile
			default:
				continue selectMove
			}

		case EAST:
			if t := currentTile.pipeType; t != HORIZONTAL && t != BEND_SOUTH_EAST && t != BEND_NORTH_EAST && t != START {
				continue selectMove
			}

			switch nextTileType {
			case HORIZONTAL, BEND_NORTH_WEST, BEND_SOUTH_WEST, START:
				return true, nextTile
			default:
				continue selectMove
			}

		case NORTH:
			if t := currentTile.pipeType; t != VERTICAL && t != BEND_NORTH_EAST && t != BEND_NORTH_WEST && t != START {
				continue
			}

			switch nextTileType {
			case VERTICAL, BEND_SOUTH_EAST, BEND_SOUTH_WEST, START:
				return true, nextTile
			default:
				continue selectMove
			}

		case WEST:
			if t := currentTile.pipeType; t != HORIZONTAL && t != BEND_NORTH_WEST && t != BEND_SOUTH_WEST && t != START {
				continue
			}

			switch nextTileType {
			case HORIZONTAL, BEND_NORTH_EAST, BEND_SOUTH_EAST, START:
				return true, nextTile
			default:
				continue selectMove
			}
		}
	}
	return false, Tile{}
}
