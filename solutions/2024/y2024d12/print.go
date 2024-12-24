package y2024d12

import (
	"fmt"

	"github.com/ddonukis/advent-of-code-golang/pkg/vec"
)

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, ch := range row {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}
func printVisited(grid [][]rune, visited map[vec.Vec2D]int) {
	for r, row := range grid {
		for c := range row {
			groupId, found := visited[vec.Vec2D{X: r, Y: c}]
			if found {
				fmt.Printf("[%2d]", groupId)
			} else {
				fmt.Print("[ ?]")
			}
		}
		fmt.Println()
	}
}
