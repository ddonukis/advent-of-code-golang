package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var linePattern regexp.Regexp = *regexp.MustCompile(`^([DRLU])\s(\d+)\s\(#([0-9a-f]{6})\)$`)

type Direction struct {
	X, Y int
}

type Point struct {
	X, Y  int
	color string
}

type Outline struct {
	points                        []Point
	leftmostPoint, rightmostPoint int
	highestPoint, lowestPoint     int
}

// move the orin to the left right corner from the starting point
func (outline *Outline) shiftOrigin() {
	for i := range outline.points {
		outline.points[i].X = outline.points[i].X - outline.leftmostPoint
		outline.points[i].Y = outline.points[i].Y - outline.highestPoint
	}
}

func (outline *Outline) MakeGrid() Grid {
	grid := make([][]*Point, outline.lowestPoint-outline.highestPoint+1)
	for rowIdx := range grid {
		grid[rowIdx] = make([]*Point, outline.rightmostPoint-outline.leftmostPoint+1)
	}
	for _, p := range outline.points {
		grid[p.Y][p.X] = &p
	}
	return grid
}

type Grid [][]*Point

func (grid Grid) String() string {
	chars := make([]rune, 0, len(grid)*(1+len(grid[0])))
	for _, row := range grid {
		for _, cell := range row {
			if cell != nil {
				chars = append(chars, '#')
			} else {
				chars = append(chars, ' ')
			}
		}
		chars = append(chars, '\n')
	}
	return string(chars)
}

func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	outlinePoints := make([]Point, 1, 10)
	outlinePoints[0] = Point{0, 0, ""}

	outline := Outline{points: outlinePoints}

	for scanner.Scan() {
		line := scanner.Text()
		direction, steps, colorCode, err := parseLine(line)
		if err != nil {
			log.Fatalf("Couldn't parse line: %v\n", err)
		}
		traceOutline(&outline, direction, steps, colorCode)
	}
	outline.shiftOrigin()
	grid := outline.MakeGrid()
	_, totalPoints := filledGrid(grid)
	fmt.Printf("P1: %d\n", totalPoints)
}

func filledGrid(grid Grid) (Grid, int) {
	updatedGrid := make([][]*Point, len(grid))
	pointsCount := 0
	for i, row := range grid {
		updatedGrid[i] = slices.Clone(row)

		intersections, consecutivePoints, connectingUp := 0, 0, 0
		for j, point := range row {
			if point != nil {
				pointsCount++
				if i-1 >= 0 && grid[i-1][j] != nil {
					connectingUp++
				}
				consecutivePoints++
			} else {
				if consecutivePoints == 1 {
					intersections++
				} else if consecutivePoints > 1 {
					intersections += connectingUp
				}

				if intersections%2 == 1 {
					updatedGrid[i][j] = &Point{X: j, Y: i}
					pointsCount++
				}
				consecutivePoints, connectingUp = 0, 0
			}
		}
	}
	return updatedGrid, pointsCount
}

func traceOutline(outline *Outline, direction Direction, steps int, color string) {
	for stepCount := 0; stepCount < steps; stepCount++ {
		prevPoint := (*outline).points[len((*outline).points)-1]
		newPoint := Point{
			X:     prevPoint.X + direction.X,
			Y:     prevPoint.Y + direction.Y,
			color: color,
		}
		(*outline).points = append((*outline).points, newPoint)
		if newPoint.X < (*outline).leftmostPoint {
			(*outline).leftmostPoint = newPoint.X
		}
		if newPoint.X > (*outline).rightmostPoint {
			(*outline).rightmostPoint = newPoint.X
		}
		if newPoint.Y < (*outline).highestPoint {
			(*outline).highestPoint = newPoint.Y
		}
		if newPoint.Y > (*outline).lowestPoint {
			(*outline).lowestPoint = newPoint.Y
		}
	}
}
func parseLine(line string) (direction Direction, steps int, colorCode string, err error) {
	matches := linePattern.FindAllStringSubmatch(line, 1)

	if len(matches) != 1 || len(matches[0]) != 4 {
		err = fmt.Errorf("line '%s' doesn't match the pattern", line)
		return
	}

	// parse direction
	switch matches[0][1] {
	case "D":
		direction = Direction{0, 1}
	case "U":
		direction = Direction{0, -1}
	case "R":
		direction = Direction{1, 0}
	case "L":
		direction = Direction{-1, 0}
	default:
		err = fmt.Errorf("invalid direction indicator '%s'", matches[0][0])
		return
	}
	// parse number of steps
	steps, err = strconv.Atoi(matches[0][2])
	colorCode = matches[0][3]
	return
}
