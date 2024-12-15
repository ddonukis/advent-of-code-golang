package y2024d15

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ddonukis/advent-of-code-golang/pkg/vec"
)

func Solve(inputPath string) {
	d1, d2 := parseInput(inputPath)

	t0 := time.Now()
	result1 := Part1(d1, d2)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	d1, d2 = parseInput(inputPath)

	t0 = time.Now()
	result2 := Part2(d1, d2)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func Part1(grid [][]Object, commands []Direction) int {
	world := World{grid: grid}

	for _, dir := range commands {
		world.MoveRobot(dir)
	}

	return world.SumGPS()
}

func Part2(grid [][]Object, commands []Direction) int {
	world := World{grid: grid}.TransformForPart2()

	fmt.Println(world)
	fmt.Println(world.robotPos)

	for _, dir := range commands {
		world.MoveRobotPart2(dir)
	}
	fmt.Println(world)
	return world.SumGPS()
}

func parseInput(inputPath string) ([][]Object, []Direction) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]Object
	var commands []Direction
	isMapSection := true
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()

		if len(line) == 0 {
			isMapSection = false
			continue
		}

		if isMapSection {
			row := parseMapRow(line)
			grid = append(grid, row)
		} else {
			dirs := parseDirections(line)
			commands = append(commands, dirs...)
		}
	}

	return grid, commands
}

func parseDirections(line string) []Direction {
	dirs := make([]Direction, len(line))
	for i, ch := range line {
		switch ch {
		case '<':
			dirs[i] = DIR_LEFT
		case '^':
			dirs[i] = DIR_UP
		case '>':
			dirs[i] = DIR_RIGHT
		case 'v':
			dirs[i] = DIR_DOWN
		default:
			msg := fmt.Sprintf("invalid direction: %q\n", ch)
			panic(msg)
		}
	}
	return dirs
}

func parseMapRow(line string) []Object {
	row := make([]Object, len(line))

	for i, ch := range line {
		switch ch {
		case '#':
			row[i] = OBJ_WALL
		case 'O':
			row[i] = OBJ_BOX
		case '.':
			row[i] = OBJ_NONE
		case '@':
			row[i] = OBJ_ROBOT
		default:
			msg := fmt.Sprintf("invalid object: %q\n", ch)
			panic(msg)
		}
	}
	return row
}

type World struct {
	grid     [][]Object
	robotPos vec.Vec2D
}

func (w World) TransformForPart2() World {
	newGrid := make([][]Object, len(w.grid))
	var newGridRobotPos vec.Vec2D

	for i, row := range w.grid {
		newRow := make([]Object, 0, 2*len(row))
		for _, obj := range row {
			switch obj {
			case OBJ_NONE:
				newRow = append(newRow, OBJ_NONE, OBJ_NONE)
			case OBJ_BOX:
				newRow = append(newRow, OBJ_BOX_L, OBJ_BOX_R)
			case OBJ_WALL:
				newRow = append(newRow, OBJ_WALL, OBJ_WALL)
			case OBJ_ROBOT:
				newGridRobotPos = vec.Vec2D{X: i, Y: len(newRow)}
				newRow = append(newRow, OBJ_ROBOT, OBJ_NONE)
			}
		}
		newGrid[i] = newRow
	}
	return World{grid: newGrid, robotPos: newGridRobotPos}
}

func (w *World) SumGPS() int {
	sum := 0
	for i, row := range w.grid {
		for j, obj := range row {
			if obj == OBJ_BOX || obj == OBJ_BOX_L {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func (w *World) MoveRobot(dir Direction) {
	curPos := w.robotPos
	if w.GetObj(curPos) != OBJ_ROBOT {
		curPos = w.LocateRobot()
		w.robotPos = curPos
	}
	var dirVec vec.Vec2D
	switch dir {
	case DIR_UP:
		dirVec = vec.Vec2D{X: -1, Y: 0}
	case DIR_LEFT:
		dirVec = vec.Vec2D{X: 0, Y: -1}
	case DIR_RIGHT:
		dirVec = vec.Vec2D{X: 0, Y: 1}
	case DIR_DOWN:
		dirVec = vec.Vec2D{X: 1, Y: 0}
	}

	var boxesOnPath []vec.Vec2D
	for {
		curPos = curPos.Add(dirVec)
		curObj := w.GetObj(curPos)
		if curObj == OBJ_BOX {
			boxesOnPath = append(boxesOnPath, curPos)
		} else if curObj == OBJ_NONE {
			// move all
			break
		} else if curObj == OBJ_WALL {
			// do nothing
			return
		}
	}
	for _, boxPos := range boxesOnPath {
		w.SetObj(boxPos, OBJ_NONE)
	}
	for _, boxPos := range boxesOnPath {
		w.SetObj(boxPos.Add(dirVec), OBJ_BOX)
	}
	w.SetObj(w.robotPos, OBJ_NONE)
	w.robotPos = w.robotPos.Add(dirVec)
	w.SetObj(w.robotPos, OBJ_ROBOT)
}

func (w *World) MoveHorizontal(dirVec vec.Vec2D) {
	// almost same as part 1
	curPos := w.robotPos
	if w.GetObj(curPos) != OBJ_ROBOT {
		curPos = w.LocateRobot()
		w.robotPos = curPos
	}

	boxes := make(map[Object][]vec.Vec2D)
	for {
		curPos = curPos.Add(dirVec)
		curObj := w.GetObj(curPos)
		if curObj == OBJ_BOX_L || curObj == OBJ_BOX_R {
			boxes[curObj] = append(boxes[curObj], curPos)
		} else if curObj == OBJ_NONE {
			break
		} else { // wall
			return
		}
	}
	for obj := range boxes {
		for _, pos := range boxes[obj] {
			w.SetObj(pos, OBJ_NONE)
		}
	}
	for obj := range boxes {
		for _, pos := range boxes[obj] {
			w.SetObj(pos.Add(dirVec), obj)
		}
	}
	w.SetObj(w.robotPos, OBJ_NONE)
	w.robotPos = w.robotPos.Add(dirVec)
	w.SetObj(w.robotPos, OBJ_ROBOT)
}

func (w *World) MoveVertical(dirVec vec.Vec2D) {
	// need to trace multiple columns
	curPos := w.robotPos
	if w.GetObj(curPos) != OBJ_ROBOT {
		curPos = w.LocateRobot()
		w.robotPos = curPos
	}

	moveQueue := make(map[Object][]vec.Vec2D)
	for {
		curPos = curPos.Add(dirVec)
		curObj := w.GetObj(curPos)
		if curObj == OBJ_NONE {
			break
		} else if curObj == OBJ_WALL {
			return
		} else if curObj == OBJ_BOX_L || curObj == OBJ_BOX_R {
			if moved := w.TryMoveObjects(curObj, curPos, dirVec, moveQueue); !moved {
				return
			} else if moved {
				break
			}
		}
	}

	for obj := range moveQueue {
		for _, pos := range moveQueue[obj] {
			w.SetObj(pos, OBJ_NONE)
		}
	}
	for obj := range moveQueue {
		for _, pos := range moveQueue[obj] {
			w.SetObj(pos.Add(dirVec), obj)
		}
	}

	w.SetObj(w.robotPos, OBJ_NONE)
	w.robotPos = w.robotPos.Add(dirVec)
	w.SetObj(w.robotPos, OBJ_ROBOT)
}

func (w *World) TryMoveObjects(
	obj Object,
	pos vec.Vec2D,
	dirVec vec.Vec2D,
	moveQueue map[Object][]vec.Vec2D,
) (success bool) {
	if obj == OBJ_NONE {
		return true
	} else if obj == OBJ_WALL {
		return false
	}

	var leftPos, rightPos, leftNextPos, rightNextPos vec.Vec2D
	if obj == OBJ_BOX_L {
		leftPos = pos
		rightPos = pos.Add(vec.Vec2D{X: 0, Y: 1})
		leftNextPos = pos.Add(dirVec)
		rightNextPos = rightPos.Add(dirVec)
	} else if obj == OBJ_BOX_R {
		leftPos = pos.Add(vec.Vec2D{X: 0, Y: -1})
		rightPos = pos
		leftNextPos = leftPos.Add(dirVec)
		rightNextPos = pos.Add(dirVec)
	}

	success = (w.TryMoveObjects(w.GetObj(leftNextPos), leftNextPos, dirVec, moveQueue) &&
		w.TryMoveObjects(w.GetObj(rightNextPos), rightNextPos, dirVec, moveQueue))

	if success {
		moveQueue[OBJ_BOX_L] = append(moveQueue[OBJ_BOX_L], leftPos)
		moveQueue[OBJ_BOX_R] = append(moveQueue[OBJ_BOX_R], rightPos)
		// fmt.Printf("Moving [: %v -> %v\n", leftPos, leftNextPos)
		// w.SetObj(leftPos, OBJ_NONE)
		// w.SetObj(rightPos, OBJ_NONE)
		// fmt.Printf("Moving ]: %v -> %v\n", rightPos, rightNextPos)
		// w.SetObj(leftNextPos, OBJ_BOX_L)
		// w.SetObj(rightNextPos, OBJ_BOX_R)
	}

	return success
}

func (w *World) MoveRobotPart2(dir Direction) {
	switch dir {
	case DIR_UP:
		w.MoveVertical(vec.Vec2D{X: -1, Y: 0})
	case DIR_DOWN:
		w.MoveVertical(vec.Vec2D{X: 1, Y: 0})
	case DIR_LEFT:
		w.MoveHorizontal(vec.Vec2D{X: 0, Y: -1})
	case DIR_RIGHT:
		w.MoveHorizontal(vec.Vec2D{X: 0, Y: 1})
	}
}

func (w *World) GetObj(pos vec.Vec2D) Object {
	return w.grid[pos.X][pos.Y]
}

func (w *World) SetObj(pos vec.Vec2D, obj Object) {
	w.grid[pos.X][pos.Y] = obj
}

func (w World) String() string {
	lines := make([]string, len(w.grid))
	for i, row := range w.grid {
		line := make([]byte, 0, len(row)+1)
		for _, obj := range row {
			switch obj {
			case OBJ_NONE:
				line = append(line, '.')
			case OBJ_WALL:
				line = append(line, '#')
			case OBJ_BOX:
				line = append(line, 'O')
			case OBJ_ROBOT:
				line = append(line, '@')
			case OBJ_BOX_L:
				line = append(line, '[')
			case OBJ_BOX_R:
				line = append(line, ']')
			}
		}
		line = append(line, '\n')
		lines[i] = string(line)
	}
	return strings.Join(lines, "")
}

func (w World) LocateRobot() vec.Vec2D {
	for i, row := range w.grid {
		for j, obj := range row {
			if obj == OBJ_ROBOT {
				return vec.Vec2D{X: i, Y: j}
			}
		}
	}
	panic("robot not found")
}

type Object int8

const (
	OBJ_NONE Object = iota
	OBJ_WALL
	OBJ_BOX
	OBJ_ROBOT
	OBJ_BOX_L // left half [ of a box []
	OBJ_BOX_R // right half ] of a box []
)

type Direction int8

const (
	DIR_UP Direction = iota
	DIR_LEFT
	DIR_RIGHT
	DIR_DOWN
)

func (d Direction) String() string {
	switch d {
	case DIR_UP:
		return "^"
	case DIR_LEFT:
		return "<"
	case DIR_RIGHT:
		return ">"
	case DIR_DOWN:
		return "V"
	default:
		return "?"
	}
}
