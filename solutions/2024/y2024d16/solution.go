package y2024d16

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ddonukis/advent-of-code-golang/pkg/vec"
)

func Solve(inputPath string) {
	data := parseInput(inputPath)

	t0 := time.Now()
	result1 := Part1(data)
	duration := time.Since(t0)
	fmt.Printf("Part 1: %d\n%d μs\n\n", result1, duration.Microseconds())

	t0 = time.Now()
	result2 := Part2(data)
	duration = time.Since(t0)
	fmt.Printf("Part 2: %d\n%d μs\n", result2, duration.Microseconds())
}

func parseInput(inputPath string) Maze {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	walls := make([][]bool, 0)
	var start, end vec.Vec2D
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()
		row := make([]bool, len(line))
		for i, ch := range line {
			if ch == '#' {
				row[i] = true
			} else if ch == 'S' {
				start = vec.Vec2D{X: len(walls), Y: i}
			} else if ch == 'E' {
				end = vec.Vec2D{X: len(walls), Y: i}
			}
		}
		walls = append(walls, row)
	}
	return Maze{Walls: walls, Start: start, StartDir: vec.Vec2D{X: 0, Y: 1}, End: end}
}

func Part1(maze Maze) int {
	fmt.Println(maze)

	mEx := MazeExplorer{maze: maze}

	minCost := mEx.Run()
	fmt.Printf("Cheapest path: %d\n", minCost)

	// pipe := make(chan int)
	// doneChan := make(chan MsgDone)
	// go func(pipe chan int, done chan MsgDone) {
	// 	for i := 0; i < 10; i++ {
	// 		pipe <- i
	// 	}
	// 	doneChan <- MsgDone{}
	// }(pipe, doneChan)

	// for {
	// 	select {
	// 	case res := <-pipe:
	// 		fmt.Println(res)
	// 	case <-doneChan:
	// 		close(pipe)
	// 		close(doneChan)
	// 		fmt.Println("Done!")
	// 		return 0
	// 	case <- time.After()
	// 	}
	// }

	// fmt.Println(maze)

	// curPos := maze.Start
	// curDir := vec.Vec2D{X: 0, Y: 1}
	// seenTiles := set.NewSet[vec.Vec2D]()
	// cost := 0
	// for {
	// 	nextPos := curDir.Add(curDir)
	// 	if !maze.isWall(nextPos) {
	// 		cost++
	// 		curPos = nextPos
	// 		continue
	// 	}
	// 	// bumped into a wall: try rotating
	// 	// if can't find a way by either rotation -> dead end
	// 	// if can go either way
	// 	nextDir1 := rotateClockwise(curDir)
	// 	nextDir2 := rotateCounterClockwise(curDir)
	// }

	return 0
}

func Part2(maze Maze) int {
	return 0
}
