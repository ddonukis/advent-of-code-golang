package y2024d12

type Vec2D struct {
	X, Y int
}

func (v Vec2D) Add(otherV Vec2D) Vec2D {
	return Vec2D{v.X + otherV.X, v.Y + otherV.Y}
}

func (v Vec2D) MulScalar(n int) Vec2D {
	return Vec2D{v.X * n, v.Y * n}
}
