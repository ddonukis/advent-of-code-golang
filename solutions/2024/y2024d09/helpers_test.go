package y2024d09

import (
	"testing"
)

func TestFind(t *testing.T) {
	mem := []Block{
		Block{0, false},
		Block{0, false},
		Block{-1, true},
		Block{-1, true},
		Block{-1, true},
		Block{0, false},
		Block{0, false},
		Block{0, false},
		Block{-1, true},
		Block{-1, true},
		Block{-1, false},
	}

	result := findFreeSpan(mem, 0, 2, len(mem))
	if result != 2 {
		t.Fatalf("expected 2 but got: %d\n", result)
	}

	result = findFreeSpan(mem, 0, 3, len(mem))
	if result != 2 {
		t.Fatalf("expected 2 but got: %d\n", result)
	}

	result = findFreeSpan(mem, 5, 2, len(mem))
	if result != 8 {
		t.Fatalf("expected 8 but got: %d\n", result)
	}

	result = findFreeSpan(mem, 8, 1, len(mem))
	if result != 8 {
		t.Fatalf("expected 8 but got: %d\n", result)
	}
}
