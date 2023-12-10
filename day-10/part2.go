package main

import "fmt"

func part2(maze Maze) {
	fmt.Println()
	fmt.Println(maze)

	insideTilesTotal := 0
	for i, row := range maze {
		for j, tileType := range row {
			if tileType == NO_PIPE { // we removed all pipe segments not part of the loop in part1
				if isInside(row[j+1:]) {
					insideTilesTotal += 1
					maze[i][j] = INSIDE
				}

			}
		}
	}

	fmt.Println()
	fmt.Println(maze)

	fmt.Printf("Tiles inside the loop: %d\n", insideTilesTotal)
}

func isInside(eastwardTiles MazeRow) bool {
	i := countIntersectionsEast(eastwardTiles)
	if i%2 == 1 {
		return true
	}
	return false
}

func countIntersectionsEast(eastwardTiles MazeRow) int {
	count := 0
	for _, tile := range eastwardTiles {
		switch tile {
		case VERTICAL:
			count++
		case BEND_NORTH_EAST, BEND_NORTH_WEST:
			count++
		case BEND_SOUTH_EAST, BEND_SOUTH_WEST:
			continue // assume our point is a little above the middle point of a cell
		default:
			continue
		}
	}
	return count
}
