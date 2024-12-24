package y2024d21

import "testing"

func TestPathToRcButton(t *testing.T) {
	//     +---+---+
	//     | ^ | A |
	// +---+---+---+
	// | < | v | > |
	// +---+---+---+
	cases := []struct {
		from    rune
		to      rune
		pathLen int
	}{
		{'A', '^', 2},
		{'A', '>', 2},
		{'A', 'v', 3},
		{'A', '<', 4},
		{'<', '>', 3},
		{'^', '>', 3},
		{'^', '<', 3},
		{'^', 'v', 2},
		{'v', '<', 2},
		{'v', '>', 2},
		{'A', 'A', 1},
		{'<', '<', 1},
	}

	for _, c := range cases {
		for i := 0; i < 2; i++ {
			var f, to rune
			if i == 0 {
				f, to = c.from, c.to
			} else {
				to, f = c.from, c.to
			}
			path := pathToRcButton(f, to)
			if len(path) != c.pathLen {
				t.Errorf("[%q] -> [%q]: %d != %d\n", f, to, len(path), c.pathLen)
			}
		}
	}
	pathToRcButton('A', '^')
}
