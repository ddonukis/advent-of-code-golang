package y2024d04

import "fmt"

type Vec2 struct {
	X int
	Y int
}

func (vec Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		vec.X + other.X,
		vec.Y + other.Y,
	}
}

func (vec Vec2) Mul(scalar int) Vec2 {
	return Vec2{
		vec.X * scalar,
		vec.Y * scalar,
	}
}

func (vec Vec2) WithinBounds(maxX, maxY int) bool {
	return vec.X >= 0 && vec.X <= maxX && vec.Y >= 0 && vec.Y <= maxY
}

type Search struct {
	startPosition Vec2
	searchWord    []byte
	cursor        int
	direction     Vec2
}

type SearchResult struct {
	startPosition Vec2
	direction     Vec2
}

func (s *Search) IsMatch(matrix CharMatrix) bool {
	for {
		nextPos := s.startPosition.Add(s.direction.Mul(s.cursor))
		if !nextPos.WithinBounds(len(matrix)-1, len(matrix[0])-1) {
			return false
		}
		if matrix[nextPos.X][nextPos.Y] != s.searchWord[s.cursor] {
			return false
		}
		s.cursor += 1
		if s.cursor == len(s.searchWord) {
			break
		}
	}
	return true
}

var directions = [8]Vec2{
	{+1, 0},
	{-1, 0},
	{0, +1},
	{0, -1},
	{+1, +1},
	{-1, -1},
	{+1, -1},
	{-1, +1},
}

func Part1(matrix CharMatrix) int {
	searches := make([]Search, 0, 100)
	for rowIdx, row := range matrix {
		for colIdx, char := range row {
			if char == 'X' {
				for _, direction := range directions {
					s := Search{
						startPosition: Vec2{rowIdx, colIdx},
						searchWord:    []byte{'X', 'M', 'A', 'S'},
						cursor:        1,
						direction:     direction,
					}
					searches = append(searches, s)
				}
			}
		}
	}
	fmt.Printf("%d searches\n", len(searches))

	found := make([]SearchResult, 0, 100)
	for _, search := range searches {
		if search.IsMatch(matrix) {
			found = append(found, SearchResult{search.startPosition, search.direction})
		}
	}

	return len(found)
}
