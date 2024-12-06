package lib

type Point2D struct {
	X int
	Y int
}

func NewPoint2D(x, y int) Point2D {
	return Point2D{
		X: x,
		Y: y,
	}
}

func (p Point2D) Add(p2 Point2D) Point2D {
	return Point2D{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

// IsInBounds assumes the map is square!
func IsInBounds[T any](relevantMap [][]T, x, y int) bool {
	return x >= 0 && x < len(relevantMap) && y >= 0 && y < len(relevantMap[0])
}

// IsPosInBounds assumes the map is square!
func IsPosInBounds[T any](relevantMap [][]T, pos Point2D) bool {
	return IsInBounds(relevantMap, pos.X, pos.Y)
}
