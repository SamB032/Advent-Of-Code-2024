package util

import "math"

type Point[T Number] struct {
	x T
	y T
}

// Calcualte the absolute distance between two points
func (point Point[T]) ManhattanDistance(other Point[T]) T {
	return T(math.Abs(float64(point.x-other.x)) + math.Abs(float64(point.y-other.y)))
}

// Calculate the Euclidean Distance from one point to another
func (point Point[T]) EuclideanDistance(other Point[T]) T {
	xPart := math.Pow(float64(point.x)-float64(other.x), 2.0)
	yPart := math.Pow(float64(point.x)-float64(other.y), 2.0)
	return T(math.Sqrt(xPart + yPart))
}
