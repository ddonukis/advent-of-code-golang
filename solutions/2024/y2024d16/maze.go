package y2024d16

import (
	"fmt"
	"math"
	"strings"

	"github.com/ddonukis/advent-of-code-golang/pkg/set"
	"github.com/ddonukis/advent-of-code-golang/pkg/vec"
)

type MazeExplorer struct {
	maze Maze
}

func (mex *MazeExplorer) Run() int {
	toExplore := make(chan NavState)

	go func() {
		toExplore <- NavState{
			Pos:  mex.maze.Start,
			Dir:  mex.maze.StartDir,
			Cost: 0,
			Seen: set.NewSet[vec.Vec2D](),
		}
	}()

	results := make(chan ExplorationResult)
	activeRunners := 0
	minCost := math.MaxInt
	// wg := &sync.WaitGroup{}
	// for i := 0; i < 4; i++ {
	// 	wg.Add(1)
	// 	go mex.Worker(toExplore, results, wg)
	// }

loop:
	for {
		select {
		case state := <-toExplore:
			// fmt.Printf("Launching runner %s %s\n", state.Pos, state.Dir)
			// if activeRunners < 8 {
			activeRunners++
			go mex.explorePath(state, toExplore, results)
			// } else {
			// 	toExplore <- state
			// }
		case r := <-results:
			activeRunners--
			fmt.Printf("(runners: %d) Reached: %t, cost: %d\n", activeRunners, r.EndReached, r.State.Cost)
			if r.EndReached && (r.State.Cost < minCost) {
				minCost = r.State.Cost
			}
			if activeRunners == 0 {
				close(results)
				close(toExplore)
				break loop
			}
		}
	}
	return minCost
}

// func (mex *MazeExplorer) Worker(
// 	toExplore chan NavState,
// 	results chan<- ExplorationResult,
// 	wg *sync.WaitGroup,
// ) {
// 	defer wg.Done()
// 	for state := range toExplore {
// 		mex.explorePath(state, toExplore, results)
// 	}
// }

func (mex *MazeExplorer) explorePath(
	state NavState,
	toExplore chan<- NavState,
	results chan<- ExplorationResult,
) {
	for {
		if state.Pos == mex.maze.End {
			results <- ExplorationResult{true, state}
			return
		}

		nextCost := 1
		var nextPos, nextDir vec.Vec2D
		pathPicked := false
		forwardPos := state.Pos.Add(state.Dir)
		if !(mex.maze.isWall(forwardPos) || state.Seen.Contains(forwardPos)) {
			pathPicked = true
			nextPos = forwardPos
			nextDir = state.Dir
		}

		leftDir := rotateCounterClockwise(state.Dir)
		leftPos := state.Pos.Add(leftDir)
		if !(mex.maze.isWall(leftPos) || state.Seen.Contains(leftPos)) {
			if !pathPicked {
				pathPicked = true
				nextPos = leftPos
				nextDir = leftDir
				nextCost = 1001
			} else {
				branchedSeen := state.Seen.Clone()
				branchedSeen.Add(leftPos)
				// fmt.Printf("branched left to %s\n", leftPos)
				toExplore <- NavState{
					Pos:  leftPos,
					Dir:  leftDir,
					Cost: state.Cost + 1001,
					Seen: branchedSeen,
				}
			}
		}

		rightDir := rotateClockwise(state.Dir)
		rightPos := state.Pos.Add(rightDir)
		if !(mex.maze.isWall(rightPos) || state.Seen.Contains(rightPos)) {
			if !pathPicked {
				pathPicked = true
				nextPos = rightPos
				nextDir = rightDir
				nextCost = 1001
			} else {
				branchedSeen := state.Seen.Clone()
				branchedSeen.Add(rightDir)
				// fmt.Printf("branched right to %s\n", rightPos)
				toExplore <- NavState{
					Pos:  rightPos,
					Dir:  rightDir,
					Cost: state.Cost + 1001,
					Seen: branchedSeen,
				}
			}
		}
		if pathPicked {
			state.Cost += nextCost
			state.Pos = nextPos
			state.Dir = nextDir
			state.Seen.Add(nextPos)
		} else {
			// dead end
			results <- ExplorationResult{false, state}
			return
		}
	}
}

type ExplorationResult struct {
	EndReached bool
	State      NavState
}

func rotateCounterClockwise(dirVec vec.Vec2D) vec.Vec2D {
	return vec.Vec2D{X: -1 * dirVec.Y, Y: dirVec.X}
}

func rotateClockwise(dirVec vec.Vec2D) vec.Vec2D {
	return vec.Vec2D{X: dirVec.Y, Y: -dirVec.X}
}

type NavState struct {
	Pos  vec.Vec2D
	Dir  vec.Vec2D
	Cost int
	Seen set.Set[vec.Vec2D]
}

type Maze struct {
	Walls    [][]bool
	Start    vec.Vec2D
	StartDir vec.Vec2D
	End      vec.Vec2D
}

func (m *Maze) isWall(position vec.Vec2D) bool {
	return m.Walls[position.X][position.Y]
}

func (m Maze) String() string {
	parts := make([]string, len(m.Walls))
	for _, row := range m.Walls {
		line := make([]rune, 0, len(row)+1)
		for _, isWall := range row {
			if isWall {
				line = append(line, '█')
				// line = append(line, '█')
			} else {
				line = append(line, ' ')
				// line = append(line, ' ')
			}
		}
		line = append(line, '\n')
		parts = append(parts, string(line))
	}
	parts = append(parts, fmt.Sprintf("Start: %s\nEnd: %s\n", m.Start, m.End))
	return strings.Join(parts, "")
}
