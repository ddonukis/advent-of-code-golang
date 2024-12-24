package vec

import "fmt"

type Vec2D struct {
	X, Y int
}

func (v Vec2D) Add(otherV Vec2D) Vec2D {
	return Vec2D{v.X + otherV.X, v.Y + otherV.Y}
}

func (v Vec2D) MulScalar(n int) Vec2D {
	return Vec2D{v.X * n, v.Y * n}
}

func (v Vec2D) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

func (v Vec2D) WithinBounds(height, width int) bool {
	return v.X >= 0 && v.Y >= 0 && v.X < height && v.Y < width
}
