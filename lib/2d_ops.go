package lib

// IsInBounds assumes the map is square!
func IsInBounds[T any](relevantMap [][]T, x, y int) bool {
	return x >= 0 && x < len(relevantMap) && y >= 0 && y < len(relevantMap[0])
}
