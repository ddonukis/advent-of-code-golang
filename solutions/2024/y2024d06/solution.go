package y2024d06

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func Solve(inputPath string) {
	tileMap, guard := parseInput(inputPath)

	t0 := time.Now()
	result1 := Part1(tileMap, guard)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(tileMap, guard)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func Part1(tileMap TileMap, guard Guard) int {
	steps := 0
	for guard.MoveOneStep(tileMap) {
		steps += 1
	}
	return guard.UniqueVisitedTiles()
}

func Part2(labMap TileMap, guard Guard) int {
	return 0
}

type TileMap [][]Tile

func (tm TileMap) String() string {
	if len(tm) == 0 {
		return ""
	}
	elements := make([]string, 0, len(tm)*(len(tm[0])+1))
	for _, row := range tm {
		for _, tile := range row {
			switch tile {
			case EMPTY:
				elements = append(elements, ".")
			case OBSTACLE:
				elements = append(elements, "#")
			case OUT_OF_BOUNDS:
				elements = append(elements, "B")
			default:
				elements = append(elements, "?")
			}
		}
		elements = append(elements, "\n")
	}
	return strings.Join(elements, "")
}

func (tm TileMap) GetTile(position Coord2D) Tile {
	return tm[position.X][position.Y]
}

func parseInput(inputPath string) (tileMap TileMap, guard Guard) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tileMap = make(TileMap, 0)
	rowCount := 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()
		row, guardPos := parseRow(line)
		rowCount += 1
		if guardPos > 0 {
			guard = NewGuard(Coord2D{rowCount, guardPos}, UP)
		}
		tileMap = append(tileMap, row)
	}

	borderRow := make([]Tile, len(tileMap[0]))
	for i := range borderRow {
		borderRow[i] = OUT_OF_BOUNDS
	}

	tileMap = slices.Insert(tileMap, 0, borderRow)
	tileMap = append(tileMap, borderRow)

	return tileMap, guard
}

// row is padded with OUT_OF_BOUNDS tiles on left and right
// guardPosition is Y coordinate of the guard, or -1 if not found
func parseRow(line string) (row []Tile, guardPosition int) {
	row = make([]Tile, len(line)+2)
	guardPosition = -1
	row[0] = OUT_OF_BOUNDS
	row[len(row)-1] = OUT_OF_BOUNDS
	for i, ch := range line {
		switch ch {
		case '.':
			row[i+1] = EMPTY
		case '#':
			row[i+1] = OBSTACLE
		case '^':
			row[i+1] = EMPTY
			guardPosition = i + 1
		default:
			err := fmt.Errorf("unknown tile type '%c' at position %d", ch, i)
			panic(err)
		}
	}
	return
}

type Coord2D struct {
	X, Y int
}

func (c Coord2D) Add(c2 Coord2D) Coord2D {
	return Coord2D{
		X: c.X + c2.X,
		Y: c.Y + c2.Y,
	}
}

type Tile uint8

const (
	OBSTACLE Tile = iota
	EMPTY
	OUT_OF_BOUNDS
)

type Direction uint8

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

var directionVecs = map[Direction]Coord2D{
	UP:    Coord2D{-1, 0},
	DOWN:  Coord2D{1, 0},
	RIGHT: Coord2D{0, 1},
	LEFT:  Coord2D{0, -1},
}

type Guard struct {
	position         Coord2D
	direction        Direction
	visitedPositions map[Coord2D]int
}

func NewGuard(position Coord2D, direction Direction) Guard {
	return Guard{
		position:         position,
		direction:        direction,
		visitedPositions: make(map[Coord2D]int),
	}
}

func (g *Guard) MoveOneStep(tileMap TileMap) (withinBounds bool) {
	dirVec := directionVecs[g.direction]
	nextPos := g.position.Add(dirVec)
	nextTile := tileMap.GetTile(nextPos)
	switch nextTile {
	case OUT_OF_BOUNDS:
		return false
	case EMPTY:
		g.position = nextPos
		g.visitedPositions[g.position] += 1
	case OBSTACLE:
		g.TurnRight()
	}
	return true
}

func (g *Guard) UniqueVisitedTiles() int {
	return len(g.visitedPositions)
}

func (g *Guard) TurnRight() {
	switch g.direction {
	case UP:
		g.direction = RIGHT
	case RIGHT:
		g.direction = DOWN
	case DOWN:
		g.direction = LEFT
	case LEFT:
		g.direction = UP
	default:
		panic("unknown direction")
	}
}
