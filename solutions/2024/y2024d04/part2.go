package y2024d04

func Part2(matrix CharMatrix) int {
	searches := make([]Search, 0, 100)
	for rowIdx, row := range matrix {
		for colIdx, char := range row {
			if char == 'M' {
				for _, direction := range directions[4:] {
					s := Search{
						startPosition: Vec2{rowIdx, colIdx},
						searchWord:    []byte{'M', 'A', 'S'},
						cursor:        1,
						direction:     direction,
					}
					searches = append(searches, s)
				}
			}
		}
	}

	xMasCount := 0
	found := make(map[Vec2]int)
	for _, search := range searches {
		if search.IsMatch(matrix) {
			midPoint := search.startPosition.Add(search.direction)
			_, exists := found[midPoint]
			if exists {
				xMasCount += 1
			}
			found[midPoint] += 1
		}
	}

	return xMasCount
}
