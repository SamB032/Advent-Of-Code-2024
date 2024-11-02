package util

import "math"

type Number interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	float32 | float64
}

type Point[T Number] struct {
	x T
	y T
}

// Calcualte the absolute distance between two points
func (point Point[T]) ManhattanDistance(other Point[T]) T {
  return T(math.Abs(float64(point.x - other.x)) + math.Abs(float64(point.y - other.y)))
}

// Calculate the Euclidean Distance from one point to another
func (point Point[T]) EuclideanDistance(other Point[T]) T {
  xPart := math.Pow(float64(point.x) - float64(other.x), 2.0)
  yPart := math.Pow(float64(point.x) - float64(other.y), 2.0)
  return T(math.Sqrt(xPart + yPart))
}
